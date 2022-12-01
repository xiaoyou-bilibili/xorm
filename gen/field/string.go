package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewFieldString(field string) String {
	return String{Field{fieldName: field}}
}

type String struct {
	Field
}

// Eq  equal to
func (str String) Eq(value string) *driver.ConditionInfo {
	return str.eq(value)
}

// Neq not equal to
func (str String) Neq(value string) *driver.ConditionInfo {
	return str.neq(value)
}

// Gt greater than
func (str String) Gt(value string) *driver.ConditionInfo {
	return str.gt(value)
}

// Gte greater or equal to
func (str String) Gte(value string) *driver.ConditionInfo {
	return str.gt(value)
}

// In ...
func (str String) In(values ...string) *driver.ConditionInfo {
	return str.in(values)
}

// NotIn ...
func (str String) NotIn(values ...string) *driver.ConditionInfo {
	return str.notIn(values)
}

// Like ...
func (str String) Like(value string) *driver.ConditionInfo {
	return str.like(value)
}

// NotLike ...
func (str String) NotLike(value string) *driver.ConditionInfo {
	return str.notLike(value)
}

func (str String) Asc() *driver.OrderInfo {
	return str.asc()
}

func (str String) Desc() *driver.OrderInfo {
	return str.desc()
}
