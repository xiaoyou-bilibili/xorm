package query

import (
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/driver"
	"testing"
)

func initMysql() driver.DbInstance {
	db, err := driver.NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
	if err != nil {
		panic(err)
	}
	return db
}

func TestPeople_Create(t *testing.T) {
	query := NewQuery(initMysql()).People
	model := &People{Id: 3, Name: "小美", Age: 14}
	fmt.Println(query.Create(model))
}

func TestPeople_Delete(t *testing.T) {
	query := NewQuery(initMysql()).People
	row, err := query.Where(query.Id.Eq(3)).Delete()
	fmt.Println(row, err)
}

func TestPeople_Update(t *testing.T) {
	query := NewQuery(initMysql()).People
	row, err := query.Where(query.Id.Eq(4)).Update(query.Name, "哈哈1")
	fmt.Println(row, err)
}

func TestPeople_Find(t *testing.T) {
	query := NewQuery(initMysql()).People
	rows, err := query.OrderBy(query.Id.Desc()).Find()
	fmt.Println(err)
	for _, row := range rows {
		fmt.Println(row)
	}
}
