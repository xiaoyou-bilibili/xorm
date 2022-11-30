package query

import (
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/driver"
	"testing"
)

func TestQuery(t *testing.T) {
	db, err := driver.NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
	if err != nil {
		panic(err)
	}

	query := NewQuery(db)

	res, err := query.People.Find()
	fmt.Println(res, err)
}
