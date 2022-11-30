package query

import (
	"github.com/xiaoyou-bilibili/xorm/gen/do"
	"github.com/xiaoyou-bilibili/xorm/gen/field"
)

type people struct {
	do.Do
	ID   field.Int64
	Name field.String
	Age  field.Int64
}

func (p people) Find() ([]*People, error) {
	return nil, nil
}
