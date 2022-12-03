package template

const QueryTemplate = `
import (
	"github.com/xiaoyou-bilibili/xorm/driver"
	"github.com/xiaoyou-bilibili/xorm/gen/do"
	"github.com/xiaoyou-bilibili/xorm/gen/field"
)

func new{{.Model}}(db driver.DbInstance) {{.Table}} {
	model := &{{.Model}}{}
	_{{.Table}} := {{.Table}}{}
	_{{.Table}}.do = do.Do{}
	// 设置数据库和表信息
	_{{.Table}}.do.SetDb(db)
	_{{.Table}}.do.SetModel(model)
	_{{.Table}}.do.SetTable(model.TableName())
	// 设置字段信息
	{{$table := .Table}}
	{{- range $i, $v := .Fields}}
	_{{$table}}.{{$v.Key}} = field.NewField{{$v.FieldType}}("{{$v.FieldName}}")
	{{- end}}

	return _{{.Table}}
}

type {{.Table}} struct {
	do   do.Do
	{{- range $i, $v := .Fields}}
	{{$v.Key}} field.{{$v.FieldType}}
	{{- end}}
}

func (p {{.Table}}) Where(conditions ...*driver.ConditionInfo) {{.Table}} {
	p.do.AddWhere(conditions...)
	return p
}

func (p {{.Table}}) Or(conditions ...*driver.ConditionInfo) {{.Table}} {
	p.do.AddOr(conditions...)
	return p
}

func (p {{.Table}}) Limit(limit int64) {{.Table}} {
	p.do.SetLimit(limit)
	return p
}

func (p {{.Table}}) Offset(offset int64) {{.Table}} {
	p.do.SetOffset(offset)
	return p
}

func (p {{.Table}}) OrderBy(orders ...*driver.OrderInfo) {{.Table}} {
	p.do.SetOrder(orders...)
	return p
}

func (p {{.Table}}) Create(models ...*{{.Model}}) error {
	return p.do.Create(models)
}

func (p {{.Table}}) Delete() (int64, error) {
	return p.do.Delete()
}

func (p {{.Table}}) Update(field field.IField, value interface{}) (int64, error) {
	return p.do.Update(field, value)
}

func (p {{.Table}}) UpdateMulti(data map[field.IField]interface{}) (int64, error) {
	return p.do.UpdateMulti(data)
}

func (p {{.Table}}) Find() ([]*{{.Model}}, error) {
	res, err := p.do.Find()
	return res.([]*{{.Model}}), err
}
`
