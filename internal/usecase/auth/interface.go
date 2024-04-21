package auth

import (
	"context"

	"configurator/internal/entity"
	entAuth "configurator/internal/entity/auth"
	"configurator/internal/usecase/internal/common"
)

type UserRepository interface {
	common.Repo

	GetUser(ctx context.Context, id entity.BigIntPK, fetchRolesDeep bool) (result entAuth.User, err error)
	ListUsers(ctx context.Context, fetchRoles bool) (result []entAuth.User, err error)
	CreateUser(ctx context.Context, user entAuth.User, assignedRoleIds ...entity.BigIntPK) (entAuth.User, error)
	DeleteUser(ctx context.Context, id entity.BigIntPK) error
	GetRolePerms(ctx context.Context, roleId entity.BigIntPK) (result []entAuth.Perm, err error)
	AuthUser(ctx context.Context, username, password string) (result entAuth.User, err error)
	IsAuth(ctx context.Context, username, password string) error

	UpdateUser(ctx context.Context, model entAuth.User) (result entAuth.User, err error)
	UpdateUserRoles(ctx context.Context, user entAuth.User, newRoleIDs ...entity.BigIntPK) (updated entAuth.User, err error)
	CreateRole(ctx context.Context, role entAuth.Role, permIds ...entity.BigIntPK) (result entAuth.Role, err error)
	ListRoles(ctx context.Context, fetchPerms bool) (result []entAuth.Role, err error)
	ListPerms(ctx context.Context) (result []entAuth.Perm, err error)
	UpdateRole(ctx context.Context, role entAuth.Role) (result entAuth.Role, err error)
	UpdateRolePerms(ctx context.Context, role entAuth.Role, newPermIds ...entity.BigIntPK) (result entAuth.Role, err error)
	DeleteRole(ctx context.Context, id entity.BigIntPK) error
}
