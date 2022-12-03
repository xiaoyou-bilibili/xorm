package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewFieldFloat64(field string) Float64 {
	return Float64{Tool{fieldName: field}}
}

type Float64 Field

func (f64 Float64) FiledName() string {
	return f64.tool.fieldName
}

// Eq  equal to
func (f64 Float64) Eq(value float64) *driver.ConditionInfo {
	return f64.tool.Eq(value)
}

// Neq not equal to
func (f64 Float64) Neq(value float64) *driver.ConditionInfo {
	return f64.tool.Neq(value)
}

// Gt greater than
func (f64 Float64) Gt(value float64) *driver.ConditionInfo {
	return f64.tool.Gt(value)
}

// Gte greater or equal to
func (f64 Float64) Gte(value float64) *driver.ConditionInfo {
	return f64.tool.Gte(value)
}

// Lt less than
func (f64 Float64) Lt(value float64) *driver.ConditionInfo {
	return f64.tool.Lt(value)
}

// Lte less or equal to
func (f64 Float64) Lte(value float64) *driver.ConditionInfo {
	return f64.tool.Lte(value)
}

// In ...
func (f64 Float64) In(values ...float64) *driver.ConditionInfo {
	return f64.tool.In(values)
}

// NotIn ...
func (f64 Float64) NotIn(values ...float64) *driver.ConditionInfo {
	return f64.tool.NotIn(values)
}

// Between ...
func (f64 Float64) Between(left float64, right float64) *driver.ConditionInfo {
	return f64.tool.Between(left, right)
}

// NotBetween ...
func (f64 Float64) NotBetween(left float64, right float64) *driver.ConditionInfo {
	return f64.tool.NotBetween(left, right)
}

func (f64 Float64) Asc() *driver.OrderInfo {
	return f64.tool.Asc()
}

func (f64 Float64) Desc() *driver.OrderInfo {
	return f64.tool.Desc()
}
