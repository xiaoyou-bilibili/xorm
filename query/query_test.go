package query

import (
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/driver"
	"github.com/xiaoyou-bilibili/xorm/utils"
	"testing"
	"time"
)

func initMysql() driver.DbInstance {
	db, err := driver.NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
	if err != nil {
		panic(err)
	}
	return db
}

func TestPeople_Create(t *testing.T) {
	query := NewQuery(initMysql()).Test
	//model := &People{Id: 3, Name: "小美", Age: 14}
	//fmt.Println(query.Create(model))
	a := "{\"age\":32}"
	model := &Test{
		Id:  2,
		Vc:  "测试2",
		Ty:  21,
		Bg:  6478576465,
		Ubg: 232363,
		Bi:  false,
		Tp:  time.Now(),
		Tx:  "ssss",
		Ft:  44.44,
		Js:  &a,
	}
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
	query := NewQuery(initMysql()).Test
	rows, err := query.OrderBy(query.Id.Desc()).Limit(1).Find()
	fmt.Println(err)
	for _, row := range rows {
		fmt.Println(utils.Interface2String(row))
	}
}
