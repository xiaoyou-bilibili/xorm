package tmp

import (
	"github.com/xiaoyou-bilibili/xorm/tmp/query"
	"testing"
)

import "gorm.io/gen"

func TestGen(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(InitDb()) // reuse your gorm db

	g.ApplyBasic(
		g.GenerateModel("people"),
	)

	// Generate the code
	g.Execute()
}

func TestSelect(t *testing.T) {
	p := query.Use(InitDb())
	//p.Transaction(func(tx *query.Query) error {
	//	return nil
	//})
	p.Person.Update(p.Person.Age, "")
	//res, err := p.Create()
	//fmt.Println(utils.Interface2String(res), err)
}
