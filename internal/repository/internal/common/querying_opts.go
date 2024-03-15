package common

import "github.com/go-pg/pg/v10/orm"

type QueryOption func(q *orm.Query) (resQ *orm.Query, wasApplied bool)

type QueryOptions []QueryOption

func (opts QueryOptions) Apply(q *orm.Query) (resQ *orm.Query, wasApplied bool) {
	_applied := false
	resQ = q
	for _, o := range opts {
		resQ, _applied = o(resQ)
		if _applied {
			wasApplied = true
		}
	}
	return
}
