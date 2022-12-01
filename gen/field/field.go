package field

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

type IField interface {
	FiledName() string
}

type Field struct {
	tool Tool
}

func (f Field) FiledName() string {
	return f.tool.fieldName
}

type Tool struct {
	fieldName string
}

func (tool Tool) Eq(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionEq,
		FieldValue: []interface{}{value},
	}
}

func (tool Tool) Neq(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionNeq,
		FieldValue: []interface{}{value},
	}
}

func (tool Tool) Gt(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionGt,
		FieldValue: []interface{}{value},
	}
}

func (tool Tool) Gte(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionGte,
		FieldValue: []interface{}{value},
	}
}

func (tool Tool) Lt(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionLt,
		FieldValue: []interface{}{value},
	}
}

func (tool Tool) Lte(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionLte,
		FieldValue: []interface{}{value},
	}
}

func (tool Tool) In(values ...interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionIn,
		FieldValue: values,
	}
}

func (tool Tool) NotIn(values ...interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionNIn,
		FieldValue: values,
	}
}

func (tool Tool) Between(left interface{}, right interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionBetween,
		FieldValue: []interface{}{left, right},
	}
}

func (tool Tool) NotBetween(left interface{}, right interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionNBetween,
		FieldValue: []interface{}{left, right},
	}
}

func (tool Tool) Like(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionLike,
		FieldValue: []interface{}{value},
	}
}

func (tool Tool) NotLike(value interface{}) *driver.ConditionInfo {
	return &driver.ConditionInfo{
		FieldName:  tool.fieldName,
		Option:     driver.ConditionOptionNLike,
		FieldValue: []interface{}{value},
	}
}

func (tool Tool) Asc() *driver.OrderInfo {
	return &driver.OrderInfo{FieldName: tool.fieldName}
}

func (tool Tool) Desc() *driver.OrderInfo {
	return &driver.OrderInfo{Desc: true, FieldName: tool.fieldName}
}
