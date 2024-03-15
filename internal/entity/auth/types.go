package auth

import (
	"github.com/Ghytro/inttest-configurator/internal/entity"
)

type (
	User struct {
		tableName struct{} `pg:"users"`

		Id entity.BigIntPK `pg:"id,pk" json:"id"`
		entity.BaseTimestamps

		UserName string `pg:"username" json:"username"`
		Password string `pg:"password" json:"-"`

		Roles []Role `pg:"many2many:user_roles,join_fk:role_id" json:"roles"`
	}

	UserRole struct {
		tableName struct{} `pg:"user_roles"`

		UserId entity.BigIntPK `pg:"user_id"`
		RoleId entity.BigIntPK `pg:"role_id"`
	}

	Role struct {
		tableName struct{} `pg:"roles"`

		Id   entity.BigIntPK `pg:"id" json:"id"`
		Name string          `pg:"name" json:"name"`
		Desc *string         `pg:"description" json:"description"`

		Perms []Perm `pg:"many2many:role_permissions,join_fk:permission_id" json:"permissions"`
	}

	RolePermission struct {
		tableName struct{} `pg:"role_permissions"`

		RoleId entity.BigIntPK `pg:"role_id"`
		PermId entity.BigIntPK `pg:"permission_id"`
	}

	Perm struct {
		tableName struct{} `pg:"permissions"`

		Id   entity.BigIntPK `pg:"id" json:"id"`
		Name EPermission     `pg:"name" json:"name"`
		Desc *string         `pg:"description" json:"description"`
	}
)
