package gen

import (
	"bytes"
	"github.com/xiaoyou-bilibili/xorm/template"
)

type Param struct {
	Table string
	Model string
	Field map[string]FieldInfo
}

type FieldInfo struct {
	FieldName string
	FieldType string
}

func queryGen(buf *bytes.Buffer, table string, field map[string]string) error {
	param := Param{
		Table: "people",
		Model: "People",
		Field: map[string]FieldInfo{
			"Id":   {FieldName: "id", FieldType: "Int64"},
			"Name": {FieldName: "name", FieldType: "String"},
			"Age":  {FieldName: "age", FieldType: "Int64"},
		},
	}
	return template.Render(template.QueryTemplate, buf, param)
}
