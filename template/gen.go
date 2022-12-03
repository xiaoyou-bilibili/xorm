package template

const GenTemplate = `
import (
	"github.com/xiaoyou-bilibili/xorm/driver"
)

func NewQuery(db driver.DbInstance) *query {
	return &Query{
		db:     db,
		{{- range $k, $v := .Tables}}
		{{$v}}: new{{$v}}(db),
		{{- end}}
	}
}

type query struct {
	db driver.DbInstance

	{{- range $k, $v := .Tables}}
	{{$v}} {{$k}}
	{{- end}}
}
`
