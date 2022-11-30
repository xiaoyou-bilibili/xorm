package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

type Int64 struct {
	Field
}

// Eq  equal to
func (i64 Int64) Eq(value int64) *driver.ConditionInfo {
	return i64.eq(value)
}

// Neq not equal to
func (i64 Int64) Neq(value int64) *driver.ConditionInfo {
	return i64.neq(value)
}

// Gt greater than
func (i64 Int64) Gt(value int64) *driver.ConditionInfo {
	return i64.gt(value)
}

// Gte greater or equal to
func (i64 Int64) Gte(value int64) *driver.ConditionInfo {
	return i64.gt(value)
}

// Lt less than
func (i64 Int64) Lt(value int64) *driver.ConditionInfo {
	return i64.lt(value)
}

// Lte less or equal to
func (i64 Int64) Lte(value int64) *driver.ConditionInfo {
	return i64.lte(value)
}

// In ...
func (i64 Int64) In(values ...int64) *driver.ConditionInfo {
	return i64.in(values)
}

// NotIn ...
func (i64 Int64) NotIn(values ...int64) *driver.ConditionInfo {
	return i64.notIn(values)
}

// Between ...
func (i64 Int64) Between(left int64, right int64) *driver.ConditionInfo {
	return i64.between(left, right)
}

// NotBetween ...
func (i64 Int64) NotBetween(left Field, right Field) *driver.ConditionInfo {
	return i64.notBetween(left, right)
}
