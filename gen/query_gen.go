package gen

import (
	"bytes"
	"github.com/xiaoyou-bilibili/xorm/template"
	"github.com/xiaoyou-bilibili/xorm/utils"
)

type Param struct {
	Table  string
	Model  string
	Fields []FieldInfo
}

type FieldInfo struct {
	Key       string
	FieldName string
	FieldType string
	IsNull    bool
}

func queryGen(buf *bytes.Buffer, table string, info []*TableField, path string) error {
	param := Param{
		Table: table,
		Model: utils.FirstUpper(table),
	}
	fields := make([]FieldInfo, 0, len(info))
	for _, field := range info {
		fields = append(fields, FieldInfo{
			Key:       utils.FirstUpper(field.FieldName),
			FieldName: field.FieldName,
			FieldType: utils.FirstUpper(field.FieldType),
			IsNull:    field.IsNull == "YES",
		})
	}
	param.Fields = fields

	if err := template.Render(template.QueryTemplate, buf, param); err != nil {
		return err
	}

	return utils.WriteGoFile(path, table+".query", buf)
}
