package gen

import (
	"bytes"
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/template"
	"github.com/xiaoyou-bilibili/xorm/utils"
	"os"
)

type queryParam struct {
	Tables map[string]string
}

func globalGen(buf *bytes.Buffer, tables []string) error {
	fields := map[string]string{}
	// 对表格转换为大小写
	for _, table := range tables {
		fields[table] = utils.FirstUpper(table)
	}
	return template.Render(template.GenTemplate, buf, queryParam{Tables: fields})
}

func writeFile(path, name string, content bytes.Buffer) error {
	return os.WriteFile(fmt.Sprintf("%s/%s.go", path, name), content.Bytes(), 0775)
}
