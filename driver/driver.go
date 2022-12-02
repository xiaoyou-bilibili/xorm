package driver

import (
	"database/sql"
	"reflect"
)

// ConditionOption 具体操作
type ConditionOption int32

const ConditionOptionLike = ConditionOption(0)      //  like
const ConditionOptionNLike = ConditionOption(1)     //  not like
const ConditionOptionEq = ConditionOption(2)        // =
const ConditionOptionNeq = ConditionOption(3)       // !=
const ConditionOptionGt = ConditionOption(4)        // >
const ConditionOptionGte = ConditionOption(5)       // >=
const ConditionOptionLt = ConditionOption(6)        // <
const ConditionOptionLte = ConditionOption(7)       // <=
const ConditionOptionIn = ConditionOption(8)        // in
const ConditionOptionNIn = ConditionOption(9)       // not in
const ConditionOptionBetween = ConditionOption(10)  // between
const ConditionOptionNBetween = ConditionOption(11) // not between

type ConditionInfo struct {
	Or         bool            // 是否为或，默认为and
	FieldName  string          // 字段名称
	Option     ConditionOption // 具体操作
	FieldValue []interface{}   // 判断的值
}

type OrderInfo struct {
	FieldName string // 待排序的字段
	Desc      bool   // 是否为降序排序
}

type FindInfo struct {
	Columns    []string         // 待查找的列
	Conditions []*ConditionInfo // 查询条件
	Orders     []*OrderInfo     // 排序条件
	Limit      *int64           // 返回限制
	Offset     *int64           // 偏移量
}

type DbInstance interface {
	// DataBaseName 获取数据库名称
	DataBaseName() string
	// SqlType 当前数据库类型
	SqlType() string
	// Create 新增数据
	Create(table string, fields map[string]interface{}) (affected int64, err error)
	// Delete 删除数据
	Delete(table string, conditions []*ConditionInfo) (affected int64, err error)
	// Update 更新数据
	Update(table string, fields map[string]interface{}, conditions []*ConditionInfo) (affected int64, err error)
	// Find 查找数据
	Find(table string, info FindInfo, p reflect.Type) (interface{}, error)
	// Transaction 事务操作
	Transaction(handle func(tx DbInstance) error) error
	// RowQuery 原始查询操作
	RowQuery(sql string, args ...interface{}) (*sql.Rows, error)
	// RowExec 原始执行操作
	RowExec(sql string, args ...interface{}) (int64, error)
}
