package do

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
	"reflect"
)

type Do struct {
	db         driver.DbInstance       // 数据库实例
	modelType  reflect.Type            // 模型的类型
	table      string                  // 表名
	columns    []string                // 待返回的行
	conditions []*driver.ConditionInfo // 查询条件
	orders     []*driver.OrderInfo     // 排序信息
	limit      *int64                  // 返回数据条数
	offset     *int64                  // 偏移量
	fields     map[string]interface{}  // 待设置的字段
}

func (do *Do) SetDb(db driver.DbInstance) {
	do.db = db
}

func (do *Do) SetModel(model interface{}) {
	do.modelType = reflect.TypeOf(model)
}

func (do *Do) SetTable(table string) {
	do.table = table
}

func (do *Do) Create(data ...interface{}) error {
	return nil
}

func (do *Do) AddWhere(infos ...*driver.ConditionInfo) {
	do.conditions = append(do.conditions, infos...)
}

func (do *Do) AddOr(infos ...*driver.ConditionInfo) {
	for _, info := range infos {
		info.Or = true
		do.conditions = append(do.conditions, info)
	}
}

func (do *Do) SetLimit(limit int64) {
	do.limit = &limit
}

func (do *Do) SetOffset(offset int64) {
	do.offset = &offset
}

func (do *Do) SetOrder(infos ...*driver.OrderInfo) {
	do.orders = append(do.orders, infos...)
}

func (do *Do) Find() (interface{}, error) {
	return do.db.Find(do.table, driver.FindInfo{
		Conditions: do.conditions,
		Limit:      do.limit,
		Offset:     do.offset,
		Orders:     do.orders,
	}, do.modelType)
}
