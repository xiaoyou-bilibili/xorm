package gen

import (
	"bytes"
	"github.com/xiaoyou-bilibili/xorm/template"
	"github.com/xiaoyou-bilibili/xorm/utils"
)

type queryParam struct {
	Tables map[string]string
}

func globalGen(buf *bytes.Buffer, tables []string, path string) error {
	fields := map[string]string{}
	// 对表格转换为大小写
	for _, table := range tables {
		fields[table] = utils.FirstUpper(table)
	}

	if err := template.Render(template.GenTemplate, buf, queryParam{Tables: fields}); err != nil {
		return err
	}

	return utils.WriteGoFile(path, "gen", buf)
}
