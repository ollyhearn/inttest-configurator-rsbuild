package auth

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"configurator/internal/entity"
	entAuth "configurator/internal/entity/auth"
	"configurator/internal/repository/internal/common"
	"configurator/pkg/database"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-pg/pg/v10/types"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type Repository struct {
	common.Mixin
}

func New(db *database.PGDB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		Mixin: common.Mixin{
			DB:          db,
			Log:         log.With(zap.String("repo", "auth_repo")),
			ErrWrapDesc: "err in auth_repo",
		},
	}
}

func (r *Repository) GetUser(ctx context.Context, id entity.BigIntPK, deepFetchRoles bool) (result entAuth.User, err error) {
	err = r.DB.RunInTransaction(ctx, func(db orm.DB) error {
		q := db.Model(&result)
		if deepFetchRoles {
			q = q.Relation("Roles")
		}
		result.Id = id
		if err := q.WherePK().Select(); err != nil {
			return err
		}
		if deepFetchRoles {
			for i, role := range result.Roles {
				p, err := r.GetRolePerms(ctx, role.Id)
				if err != nil {
					return err
				}
				result.Roles[i].Perms = p
			}
		}
		return nil
	})
	if err != nil {
		return entAuth.User{}, r.WrapErr(err)
	}
	return result, nil
}

func (r *Repository) ListUsers(ctx context.Context, fetchRoles bool) (result []entAuth.User, err error) {
	err = r.DB.RunInTransaction(ctx, func(db orm.DB) error {
		q := db.Model(&result)
		if fetchRoles {
			q.Relation("Roles")
		}
		return q.Select()
	})
	if err != nil {
		return nil, r.WrapErr(err)
	}
	return result, nil
}

func (r *Repository) CreateUser(ctx context.Context, user entAuth.User, assignedRoles ...string) (entAuth.User, error) {
	const errDesc = "ошибка создания пользователя"

	assignedRoles = lo.Uniq(assignedRoles)

	err := r.DB.RunInTransaction(ctx, func(db orm.DB) error {
		if err := r.createUserPreconds(ctx, user, assignedRoles...); err != nil {
			return errors.Wrap(err, errDesc)
		}

		user.CreatedAt = time.Now()
		user.UpdatedAt = types.NullTime{}
		user.DeletedAt = types.NullTime{}
		if _, err := db.Model(&user).Returning("*").Insert(); err != nil {
			r.Log.Error(err)
			return errors.New("попробуйте позже")
		}
		if _, err := db.Model(&user).
			Set("password=crypt(?, gen_salt('bf'))", user.Password).
			WherePK().
			Update(); err != nil {
			r.Log.Error(err)
			return errors.New("попробуйте позже")
		}

		if len(assignedRoles) > 0 {
			var roles []*entAuth.Role
			if err := db.Model(&roles).WhereIn("name IN (?)", assignedRoles).Select(); err != nil {
				r.Log.Error(err)
				return errors.New("получение списка ролей")
			}
			userRoles := lo.Map(roles, func(r *entAuth.Role, _ int) *entAuth.UserRole {
				return &entAuth.UserRole{
					UserId: user.Id,
					RoleId: r.Id,
				}
			})
			if _, err := db.Model(&userRoles).Insert(); err != nil {
				r.Log.Error(err)
				return errors.New("ошибка записи ролей")
			}
		}
		return nil
	})
	if err != nil {
		return entAuth.User{}, errors.Wrap(err, "ошибка создания пользователя")
	}
	return user, nil
}

func (r *Repository) IsAuth(ctx context.Context, username, password string) error {
	return r.DB.RunInTransaction(ctx, func(db orm.DB) error {
		c, err := db.Model((*entAuth.User)(nil)).
			Where("username=? AND password=crypt(?, password)", username, password).
			Count()
		if err != nil {
			return err
		}
		if c == 0 {
			return errors.New("неправильный логин или пароль")
		}
		return nil
	})
}

func (r *Repository) AuthUser(ctx context.Context, username, password string) (result entAuth.User, err error) {
	err = r.DB.RunInTransaction(ctx, func(db orm.DB) error {
		err := db.Model(&result).
			Where("username=? AND password=crypt(?, password)", username, password).
			Select()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return entAuth.User{}, nil
	}
	return result, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id entity.BigIntPK) error {
	return r.DB.RunInTransaction(ctx, func(db orm.DB) error {
		if _, err := db.Exec("DELETE FROM users WHERE id=?", id); err != nil {
			return err
		}
		return nil
	})
}

func (r *Repository) createUserPreconds(ctx context.Context, user entAuth.User, assignedRoles ...string) error {
	return r.DB.RunInTransaction(ctx, func(db orm.DB) error {
		var users []*entAuth.User
		if err := db.Model(&users).Where("username=?", user.UserName).Select(); err != nil {
			return err
		}
		if len(users) != 0 {
			return errors.New("имя пользователя обязано быть уникальным в системе")
		}

		if len(assignedRoles) > 0 {
			var rolesCount int
			if _, err := db.Query(
				pg.Scan(&rolesCount),
				"SELECT COUNT(*) FROM roles WHERE name IN (?)",
				pg.In(assignedRoles),
			); err != nil {
				r.Log.Error(err)
				return errors.New("получение списка ролей")
			}
			if rolesCount != len(assignedRoles) {
				return errors.New("имеются некорректные значения назначаемых ролей")
			}
		}
		return nil
	})
}

func (r *Repository) GetRolePerms(ctx context.Context, roleId entity.BigIntPK) (result []entAuth.Perm, err error) {
	err = r.DB.RunInTransaction(ctx, func(db orm.DB) error {
		role := entAuth.Role{
			Id: roleId,
		}
		if err := db.Model(&role).Relation("Perms").WherePK().Select(); err != nil {
			return err
		}
		result = role.Perms
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
