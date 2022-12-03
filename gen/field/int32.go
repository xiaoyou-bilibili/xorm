package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewFieldInt32(field string) Int32 {
	return Int32{Tool{fieldName: field}}
}

type Int32 Field

func (i32 Int32) FiledName() string {
	return i32.tool.fieldName
}

// Eq  equal to
func (i32 Int32) Eq(value int32) *driver.ConditionInfo {
	return i32.tool.Eq(value)
}

// Neq not equal to
func (i32 Int32) Neq(value int32) *driver.ConditionInfo {
	return i32.tool.Neq(value)
}

// Gt greater than
func (i32 Int32) Gt(value int32) *driver.ConditionInfo {
	return i32.tool.Gt(value)
}

// Gte greater or equal to
func (i32 Int32) Gte(value int32) *driver.ConditionInfo {
	return i32.tool.Gte(value)
}

// Lt less than
func (i32 Int32) Lt(value int32) *driver.ConditionInfo {
	return i32.tool.Lt(value)
}

// Lte less or equal to
func (i32 Int32) Lte(value int32) *driver.ConditionInfo {
	return i32.tool.Lte(value)
}

// In ...
func (i32 Int32) In(values ...int32) *driver.ConditionInfo {
	return i32.tool.In(values)
}

// NotIn ...
func (i32 Int32) NotIn(values ...int32) *driver.ConditionInfo {
	return i32.tool.NotIn(values)
}

// Between ...
func (i32 Int32) Between(left int32, right int32) *driver.ConditionInfo {
	return i32.tool.Between(left, right)
}

// NotBetween ...
func (i32 Int32) NotBetween(left int32, right int32) *driver.ConditionInfo {
	return i32.tool.NotBetween(left, right)
}

func (i32 Int32) Asc() *driver.OrderInfo {
	return i32.tool.Asc()
}

func (i32 Int32) Desc() *driver.OrderInfo {
	return i32.tool.Desc()
}
