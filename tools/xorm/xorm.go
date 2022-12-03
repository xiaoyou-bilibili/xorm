package main

import (
	"flag"
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/driver"
	"github.com/xiaoyou-bilibili/xorm/gen"
	"strconv"
	"strings"
)

type cmdConf struct {
	IP       string
	Port     int
	User     string
	Passwd   string
	Database string
	Table    string
	Path     string
	Pkg      string
}

func argParse() *cmdConf {
	ip := flag.String("ip", "localhost", "数据库ip")
	port := flag.String("port", "3306", "数据库端口")
	user := flag.String("user", "root", "数据库用户名")
	passwd := flag.String("passwd", "", "数据库密码")
	database := flag.String("db", "", "数据库名")
	table := flag.String("table", "", "表名,多个表用,隔开")
	path := flag.String("path", "./query", "生成代码的路径")
	pkg := flag.String("pkg", "query", "包名")
	flag.Parse()
	pt, err := strconv.Atoi(*port)
	if err != nil {
		panic("端口必须为数字")
	}
	if *database == "" {
		panic("数据库名不能为空")
	}
	if *table == "" {
		panic("表名不能为空")
	}

	var cmd cmdConf
	cmd.IP = *ip
	cmd.Port = pt
	cmd.User = *user
	cmd.Passwd = *passwd
	cmd.Database = *database
	cmd.Table = *table
	cmd.Path = *path
	cmd.Pkg = *pkg
	return &cmd
}

//go:generate go run main.go -ip 192.168.1.10 -port 8006 -user root -passwd xiaoyou -db demo -table people -path ./tmp -pkg tmp
func main() {
	cmd := argParse()
	// 直接调用函数来生成
	db, err := driver.NewMysqlDevice(cmd.IP, cmd.Port, cmd.User, cmd.Passwd, cmd.Database)
	if err != nil {
		panic(err)
	}
	err = gen.GenerateCode(db, strings.Split(cmd.Table, ","), gen.GenerateConfig{
		Path: cmd.Path,
		Pkg:  cmd.Pkg,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("代码生成完毕！(*^▽^*)")
}
