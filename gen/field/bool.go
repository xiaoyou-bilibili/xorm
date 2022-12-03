package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewFieldBool(field string) Bool {
	return Bool{Tool{fieldName: field}}
}

type Bool Field

func (bo Bool) FiledName() string {
	return bo.tool.fieldName
}

// Eq  equal to
func (bo Bool) Eq(value bool) *driver.ConditionInfo {
	return bo.tool.Eq(value)
}

// Neq not equal to
func (bo Bool) Neq(value bool) *driver.ConditionInfo {
	return bo.tool.Neq(value)
}

// In ...
func (bo Bool) In(values ...bool) *driver.ConditionInfo {
	return bo.tool.In(values)
}

// NotIn ...
func (bo Bool) NotIn(values ...bool) *driver.ConditionInfo {
	return bo.tool.NotIn(values)
}

func (bo Bool) Asc() *driver.OrderInfo {
	return bo.tool.Asc()
}

func (bo Bool) Desc() *driver.OrderInfo {
	return bo.tool.Desc()
}
