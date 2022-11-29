package driver

import (
	"fmt"
	"testing"
)

type People struct {
	Id   int64  `xorm:"id"`
	Name string `xorm:"name"`
	Age  int64  `xorm:"age"`
}

func TestAddStr(t *testing.T) {
	db, err := NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
	if err != nil {
		panic(err)
	}
	row, err := db.Create("people", map[string]interface{}{
		"id":   5,
		"name": "小游",
		"age":  10,
	})
	fmt.Println(row, err)
	//// See "Important settings" section.
	//
	//rows, err := db.Query("select * from people;")
	//var res []*People
	//err = utils.ConvertRows2Struct(rows, &res)
	//fmt.Println("结果", utils.Interface2String(res))
	//fmt.Println(err)
}
