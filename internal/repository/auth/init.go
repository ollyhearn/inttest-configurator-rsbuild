package auth

import (
	entAuth "github.com/Ghytro/inttest-configurator/internal/entity/auth"
	"github.com/go-pg/pg/v10/orm"
)

func init() {
	orm.RegisterTable((*entAuth.UserRole)(nil))
	orm.RegisterTable((*entAuth.RolePermission)(nil))
}
