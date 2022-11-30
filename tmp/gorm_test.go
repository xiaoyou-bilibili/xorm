package tmp

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type People struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func InitDb() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:xiaoyou@tcp(192.168.1.10:8006)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	return db
}

func TestGet(t *testing.T) {
	db := InitDb()
	people := &People{}
	db.Table("people").Find(people)
	fmt.Println(people)
}
