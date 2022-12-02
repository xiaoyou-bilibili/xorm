package template

const GenTemplate = `
import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewQuery(db driver.DbInstance) *Query {
	return &Query{
		db:     db,
		{{- range $k, $v := .Tables}}
		{{$v}}: New{{$v}}(db),
		{{- end}}
	}
}

type Query struct {
	db driver.DbInstance

	{{- range $k, $v := .Tables}}
	{{$v}} {{$k}}
	{{- end}}
}
`
