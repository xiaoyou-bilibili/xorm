package gen

import (
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/driver"
	"github.com/xiaoyou-bilibili/xorm/utils"
	"testing"
)

func initMysql() driver.DbInstance {
	db, err := driver.NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
	if err != nil {
		panic(err)
	}
	return db
}

func TestGetTable(t *testing.T) {
	db := initMysql()
	res, err := getTableFields(db, "people")
	fmt.Println(utils.Interface2String(res), err)
}

func TestGenerateCode(t *testing.T) {
	db := initMysql()
	err := GenerateCode(db, []string{"people"}, GenerateConfig{Path: "../query", Pkg: "query"})
	fmt.Println(err)
}
