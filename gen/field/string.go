package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewFieldString(field string) String {
	return String{Tool{fieldName: field}}
}

type String Field

func (str String) FiledName() string {
	return str.tool.fieldName
}

// Eq  equal to
func (str String) Eq(value string) *driver.ConditionInfo {
	return str.tool.Eq(value)
}

// Neq not equal to
func (str String) Neq(value string) *driver.ConditionInfo {
	return str.tool.Neq(value)
}

// Gt greater than
func (str String) Gt(value string) *driver.ConditionInfo {
	return str.tool.Gt(value)
}

// Gte greater or equal to
func (str String) Gte(value string) *driver.ConditionInfo {
	return str.tool.Gte(value)
}

// In ...
func (str String) In(values ...string) *driver.ConditionInfo {
	return str.tool.In(values)
}

// NotIn ...
func (str String) NotIn(values ...string) *driver.ConditionInfo {
	return str.tool.NotIn(values)
}

// Like ...
func (str String) Like(value string) *driver.ConditionInfo {
	return str.tool.Like(value)
}

// NotLike ...
func (str String) NotLike(value string) *driver.ConditionInfo {
	return str.tool.NotLike(value)
}

func (str String) Asc() *driver.OrderInfo {
	return str.tool.Asc()
}

func (str String) Desc() *driver.OrderInfo {
	return str.tool.Desc()
}
