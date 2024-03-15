package auth

import (
	"context"

	"github.com/Ghytro/inttest-configurator/internal/entity"
	entAuth "github.com/Ghytro/inttest-configurator/internal/entity/auth"
	"github.com/Ghytro/inttest-configurator/internal/usecase/internal/common"
)

type UserRepository interface {
	common.Repo

	GetUser(ctx context.Context, id entity.BigIntPK, fetchRolesDeep bool) (result entAuth.User, err error)
	ListUsers(ctx context.Context, fetchRoles bool) (result []entAuth.User, err error)
	CreateUser(ctx context.Context, user entAuth.User, assignedRoles ...string) (entAuth.User, error)
	DeleteUser(ctx context.Context, id entity.BigIntPK) error
	GetRolePerms(ctx context.Context, roleId entity.BigIntPK) (result []entAuth.Perm, err error)
	AuthUser(ctx context.Context, username, password string) (result entAuth.User, err error)
	IsAuth(ctx context.Context, username, password string) error
}
