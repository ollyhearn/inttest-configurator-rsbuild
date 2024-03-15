package entity

import (
	"strconv"
	"time"

	"github.com/go-pg/pg/v10/types"
)

type BigIntPK uint64

func ParseBigIntPK(s string) (BigIntPK, error) {
	result, err := strconv.ParseUint(s, 10, 64)
	return BigIntPK(result), err
}

type BaseTimestamps struct {
	CreatedAt time.Time      `pg:"created_at"`
	UpdatedAt types.NullTime `pg:"updated_at"`
	DeletedAt types.NullTime `pg:"deleted_at"`
}
