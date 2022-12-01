package driver

import (
	"errors"
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
	db.Transaction(func(tx DbInstance) error {
		tx.Create("people", map[string]interface{}{
			"id":   5,
			"name": "小游",
			"age":  10,
		})
		//return nil
		return errors.New("123")
	})
	//fmt.Println(db.Delete("people", []*ConditionInfo{
	//	{Option: ConditionOptionNIn, FieldName: "name", FieldValue: []interface{}{1, 2}},
	//}))
	//fmt.Println(db.Update("people", map[string]interface{}{
	//	"name": "小游",
	//}, []*ConditionInfo{
	//	{Option: ConditionOptionEq, FieldName: "id", FieldValue: []interface{}{1}},
	//}))
	//limit := int64(tx.Commit())
	//res, err := db.Find("people", FindInfo{
	//	Conditions: []*ConditionInfo{{Option: ConditionOptionIn, FieldValue: []interface{}{1, 2}, FieldName: "id"}},
	//	Orders:     []*OrderInfo{{FieldName: "id", Desc: true}},
	//	Limit:      &limit,
	//}, reflect.TypeOf(People{}))
	//fmt.Println("结果", res)
	//res2 := utils.Interface2String(res)
	//fmt.Println(res2)
}
