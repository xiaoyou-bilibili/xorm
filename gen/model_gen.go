package gen

import (
	"bytes"
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/template"
	"github.com/xiaoyou-bilibili/xorm/utils"
)

type Model interface {
	TableName() string
}

func modelGen(buf *bytes.Buffer, table string, info []*TableField, path string) error {
	param := Param{
		Table: table,
		Model: utils.FirstUpper(table),
	}
	fields := make([]FieldInfo, 0, len(info))
	for _, field := range info {
		fields = append(fields, FieldInfo{
			Key:       utils.FirstUpper(field.FieldName),
			FieldName: fmt.Sprintf("`xorm:\"%s\" json:\"%s\"`", field.FieldName, field.FieldName),
			FieldType: field.FieldType,
		})
	}
	param.Fields = fields
	if err := template.Render(template.ModelTemplate, buf, param); err != nil {
		return err
	}

	return writeFile(path, table+".model", buf)
}
