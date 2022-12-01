package query

import (
	"github.com/xiaoyou-bilibili/xorm/driver"
	"github.com/xiaoyou-bilibili/xorm/gen/do"
	"github.com/xiaoyou-bilibili/xorm/gen/field"
)

func NewPeople(db driver.DbInstance) people {
	model := &People{}
	_people := people{}
	_people.do = do.Do{}
	// 设置数据库和表信息
	_people.do.SetDb(db)
	_people.do.SetModel(model)
	_people.do.SetTable(model.TableName())
	// 设置字段信息
	_people.ID = field.NewFieldInt64("id")
	_people.Name = field.NewFieldString("name")
	_people.Age = field.NewFieldInt64("age")

	return _people
}

type people struct {
	do   do.Do
	ID   field.Int64
	Name field.String
	Age  field.Int64
}

func (p people) Where(conditions ...*driver.ConditionInfo) people {
	p.do.AddWhere(conditions...)
	return p
}

func (p people) Or(conditions ...*driver.ConditionInfo) people {
	p.do.AddOr(conditions...)
	return p
}

func (p people) Limit(limit int64) people {
	p.do.SetLimit(limit)
	return p
}

func (p people) Offset(offset int64) people {
	p.do.SetOffset(offset)
	return p
}

func (p people) OrderBy(orders ...*driver.OrderInfo) people {
	p.do.SetOrder(orders...)
	return p
}

func (p people) Create(models ...*People) error {
	return p.do.Create(models)
}

func (p people) Delete() (int64, error) {
	return p.do.Delete()
}

func (p people) Update(field field.IField, value interface{}) (int64, error) {
	return p.do.Update(field, value)
}

func (p people) UpdateMulti(data map[field.IField]interface{}) (int64, error) {
	return p.do.UpdateMulti(data)
}

func (p people) Find() ([]*People, error) {
	res, err := p.do.Find()
	return res.([]*People), err
}