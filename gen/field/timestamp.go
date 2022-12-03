package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
	"time"
)

func NewFieldTimestamp(field string) Timestamp {
	return Timestamp{Tool{fieldName: field}}
}

type Timestamp Field

func (tp Timestamp) FiledName() string {
	return tp.tool.fieldName
}

// Eq  equal to
func (tp Timestamp) Eq(value time.Time) *driver.ConditionInfo {
	return tp.tool.Eq(value)
}

// Neq not equal to
func (tp Timestamp) Neq(value time.Time) *driver.ConditionInfo {
	return tp.tool.Neq(value)
}

// Gt greater than
func (tp Timestamp) Gt(value time.Time) *driver.ConditionInfo {
	return tp.tool.Gt(value)
}

// Gte greater or equal to
func (tp Timestamp) Gte(value time.Time) *driver.ConditionInfo {
	return tp.tool.Gte(value)
}

// Lt less than
func (tp Timestamp) Lt(value time.Time) *driver.ConditionInfo {
	return tp.tool.Lt(value)
}

// Lte less or equal to
func (tp Timestamp) Lte(value time.Time) *driver.ConditionInfo {
	return tp.tool.Lte(value)
}

// In ...
func (tp Timestamp) In(values ...time.Time) *driver.ConditionInfo {
	return tp.tool.In(values)
}

// NotIn ...
func (tp Timestamp) NotIn(values ...time.Time) *driver.ConditionInfo {
	return tp.tool.NotIn(values)
}

// Between ...
func (tp Timestamp) Between(left time.Time, right time.Time) *driver.ConditionInfo {
	return tp.tool.Between(left, right)
}

// NotBetween ...
func (tp Timestamp) NotBetween(left time.Time, right time.Time) *driver.ConditionInfo {
	return tp.tool.NotBetween(left, right)
}

func (tp Timestamp) Asc() *driver.OrderInfo {
	return tp.tool.Asc()
}

func (tp Timestamp) Desc() *driver.OrderInfo {
	return tp.tool.Desc()
}
