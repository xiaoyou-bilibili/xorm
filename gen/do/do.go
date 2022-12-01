package do

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
	"github.com/xiaoyou-bilibili/xorm/gen/field"
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

func (do *Do) Create(data interface{}) error {
	// 待插入的数据
	var insertData []map[string]interface{}
	// 遍历切片
	vf := reflect.ValueOf(data)
	for i := 0; i < vf.Len(); i++ {
		// 获取切片对于的值和类型
		item := vf.Index(i).Elem()
		itemType := reflect.TypeOf(item.Interface())
		data := map[string]interface{}{}
		// 遍历结构体每个字段，并自动添加到map中
		for j := 0; j < itemType.NumField(); j++ {
			data[itemType.Field(j).Tag.Get("xorm")] = item.Field(j).Interface()
		}
		insertData = append(insertData, data)
	}
	// 使用事务来进行操作
	return do.db.Transaction(func(tx driver.DbInstance) error {
		for _, insert := range insertData {
			if _, err := tx.Create(do.table, insert); err != nil {
				return err
			}
		}
		return nil
	})
}

func (do *Do) Delete() (int64, error) {
	return do.db.Delete(do.table, do.conditions)
}

// Update 更新单个字段
func (do *Do) Update(field field.IField, value interface{}) (int64, error) {
	return do.db.Update(do.table, map[string]interface{}{field.FiledName(): value}, do.conditions)
}

// UpdateMulti 更新多个字段
func (do *Do) UpdateMulti(info map[field.IField]interface{}) (int64, error) {
	data := make(map[string]interface{})
	for fd, value := range info {
		data[fd.FiledName()] = value
	}
	return do.db.Update(do.table, data, do.conditions)
}

func (do *Do) Find() (interface{}, error) {
	return do.db.Find(do.table, driver.FindInfo{
		Conditions: do.conditions,
		Limit:      do.limit,
		Offset:     do.offset,
		Orders:     do.orders,
	}, do.modelType)
}
