package driver

import (
	"errors"
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/utils"
	"reflect"
	"testing"
)

type People struct {
	Id   int64  `xorm:"id"`
	Name string `xorm:"name"`
	Age  int64  `xorm:"age"`
}

func initMysql() DbInstance {
	db, err := NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
	if err != nil {
		panic(err)
	}
	return db
}

func TestMysqlDriver_Create(t *testing.T) {
	db := initMysql()
	res, err := db.Create("people", map[string]interface{}{
		"id":   4,
		"name": "测试",
		"age":  20,
	})
	fmt.Println(res, err)
}

func TestMysqlDriver_Delete(t *testing.T) {
	db := initMysql()
	fmt.Println(db.Delete("people", []*ConditionInfo{
		{Option: ConditionOptionNIn, FieldName: "name", FieldValue: []interface{}{1, 2}},
	}))
}

func TestMysqlDriver_Update(t *testing.T) {
	db := initMysql()
	fmt.Println(db.Update("people", map[string]interface{}{
		"name": "小游",
	}, []*ConditionInfo{
		{Option: ConditionOptionEq, FieldName: "id", FieldValue: []interface{}{1}},
	}))
}

func TestMysqlDriver_Find(t *testing.T) {
	db := initMysql()
	limit := int64(1)
	res, err := db.Find("people", FindInfo{
		Conditions: []*ConditionInfo{{Option: ConditionOptionIn, FieldValue: []interface{}{1, 2}, FieldName: "id"}},
		Orders:     []*OrderInfo{{FieldName: "id", Desc: true}},
		Limit:      &limit,
	}, reflect.TypeOf(People{}))
	fmt.Println("结果", res, err)
	res2 := utils.Interface2String(res)
	fmt.Println(res2)
}

func TestTransaction(t *testing.T) {
	db := initMysql()
	db.Transaction(func(tx DbInstance) error {
		tx.Create("people", map[string]interface{}{
			"id":   5,
			"name": "小游",
			"age":  10,
		})
		//return nil
		return errors.New("123")
	})
}
