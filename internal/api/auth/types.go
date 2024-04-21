package auth

import (
	"time"

	"configurator/internal/entity"
)

type (
	createUserRequest struct {
		UserName string            `json:"username"`
		Password string            `json:"password"`
		Roles    []entity.BigIntPK `json:"role_ids"`
	}
	createUserResponse struct {
		Id        entity.BigIntPK `json:"id"`
		CreatedAt time.Time       `json:"created_at"`
	}

	listUsersResponseItem struct {
		Id        entity.BigIntPK   `json:"id"`
		UserName  string            `json:"username"`
		CreatedAt time.Time         `json:"created_at"`
		Roles     []entity.BigIntPK `json:"role_ids"`
	}

	updateUserRequest struct {
		UserName string            `json:"username"`
		Roles    []entity.BigIntPK `json:"role_ids"`
	}

	authRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	authResponse struct {
		Token string `json:"token"`
	}

	roleCreateRequest struct {
		Name    string            `json:"name"`
		Desc    *string           `json:"desc"`
		PermIds []entity.BigIntPK `json:"perm_ids"`
	}
	roleCreateResponse struct {
		Id entity.BigIntPK `json:"id"`
	}

	listRoleResponseItem struct {
		Id      entity.BigIntPK   `json:"id"`
		Name    string            `json:"name"`
		Desc    *string           `json:"desc"`
		PermIds []entity.BigIntPK `json:"perm_ids"`
	}

	listPermResponseItem struct {
		Id   entity.BigIntPK `json:"id"`
		Name string          `json:"name"`
		Desc *string         `json:"desc"`
	}

	updateRoleRequest struct {
		Name    string            `json:"name"`
		Desc    *string           `json:"desc"`
		PermIds []entity.BigIntPK `json:"perm_ids"`
	}
	updateRoleResponse struct {
		// updated_at, row_index...
	}
)
