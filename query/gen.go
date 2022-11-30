package query

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewQuery(db driver.DbInstance) *Query {
	return &Query{db: db}
}

type Query struct {
	db driver.DbInstance

	People people
}