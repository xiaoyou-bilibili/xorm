package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

type Field struct {
	fieldName string
}

func (field Field) eq(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionEq,
		FieldValue: []interface{}{value},
	}
}

func (field Field) neq(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionNeq,
		FieldValue: []interface{}{value},
	}
}

func (field Field) gt(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionGt,
		FieldValue: []interface{}{value},
	}
}

func (field Field) gte(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionGte,
		FieldValue: []interface{}{value},
	}
}

func (field Field) lt(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionLt,
		FieldValue: []interface{}{value},
	}
}

func (field Field) lte(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionLte,
		FieldValue: []interface{}{value},
	}
}

func (field Field) in(values ...interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionIn,
		FieldValue: values,
	}
}

func (field Field) notIn(values ...interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionNIn,
		FieldValue: values,
	}
}

func (field Field) between(left interface{}, right interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionBetween,
		FieldValue: []interface{}{left, right},
	}
}

func (field Field) notBetween(left interface{}, right interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionNBetween,
		FieldValue: []interface{}{left, right},
	}
}

func (field Field) like(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionLike,
		FieldValue: []interface{}{value},
	}
}

func (field Field) notLike(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  field.fieldName,
		Option:     driver.ConditionOptionNLike,
		FieldValue: []interface{}{value},
	}
}

func (field Field) asc() *driver.OrderInfo {
	return &driver.OrderInfo{FieldName: field.fieldName}
}

func (field Field) desc() *driver.OrderInfo {
	return &driver.OrderInfo{Desc: true, FieldName: field.fieldName}
}
