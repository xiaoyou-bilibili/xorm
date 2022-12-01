package query

import (
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/driver"
	"github.com/xiaoyou-bilibili/xorm/gen/field"
	"testing"
)

func TestQuery(t *testing.T) {
	db, err := driver.NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
	if err != nil {
		panic(err)
	}

	query := NewQuery(db).People
	//data := &People{
	//	Id:   3,
	//	Name: "测速",
	//	Age:  44,
	//}
	//err = query.Create(data)
	//fmt.Println(err)
	fmt.Println(query.Where(query.ID.Eq(2)).UpdateMulti(map[field.IField]interface{}{
		query.Name: "测试11",
		query.Age:  33,
	}))
	//for _, people := range res {
	//	fmt.Println("结果", people)
	//}
}
