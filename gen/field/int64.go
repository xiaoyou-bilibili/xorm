package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewFieldInt64(field string) Int64 {
	return Int64{Tool{fieldName: field}}
}

type Int64 Field

func (i64 Int64) FiledName() string {
	return i64.tool.fieldName
}

// Eq  equal to
func (i64 Int64) Eq(value int64) *driver.ConditionInfo {
	return i64.tool.Eq(value)
}

// Neq not equal to
func (i64 Int64) Neq(value int64) *driver.ConditionInfo {
	return i64.tool.Neq(value)
}

// Gt greater than
func (i64 Int64) Gt(value int64) *driver.ConditionInfo {
	return i64.tool.Gt(value)
}

// Gte greater or equal to
func (i64 Int64) Gte(value int64) *driver.ConditionInfo {
	return i64.tool.Gte(value)
}

// Lt less than
func (i64 Int64) Lt(value int64) *driver.ConditionInfo {
	return i64.tool.Lt(value)
}

// Lte less or equal to
func (i64 Int64) Lte(value int64) *driver.ConditionInfo {
	return i64.tool.Lte(value)
}

// In ...
func (i64 Int64) In(values ...int64) *driver.ConditionInfo {
	return i64.tool.In(values)
}

// NotIn ...
func (i64 Int64) NotIn(values ...int64) *driver.ConditionInfo {
	return i64.tool.NotIn(values)
}

// Between ...
func (i64 Int64) Between(left int64, right int64) *driver.ConditionInfo {
	return i64.tool.Between(left, right)
}

// NotBetween ...
func (i64 Int64) NotBetween(left int64, right int64) *driver.ConditionInfo {
	return i64.tool.NotBetween(left, right)
}

func (i64 Int64) Asc() *driver.OrderInfo {
	return i64.tool.Asc()
}

func (i64 Int64) Desc() *driver.OrderInfo {
	return i64.tool.Desc()
}
