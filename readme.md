# XORM
> 一款简易的orm框架，支持多数据库源，自动模型生成，自动代码生成等功能（此项目只是对go的一些探索，本人不打算维护）

## 功能特点
- 完全不依赖其他开源库（驱动除外）
- 支持多数据源（目前只有mysql）
- 结构和功能非常简单（适合学习orm框架的思想）
- 支持CRUD和事务操作
- 支持把查询结果映射到struct中
- 支持代码生成，自动为现有数据表生成CRUD代码

## 代码生成

代码生成有两钟方法，第一种就是直接在项目中集成代码，第二种就是使用命令行工具一键生成

### 项目集成
先安装一下依赖
```shell
go get github.com/xiaoyou-bilibili/xorm@latest
```
生成代码如下，注意参数替换
```go
package main

import (
	"fmt"
	"github.com/xiaoyou-bilibili/xorm/driver"
	"github.com/xiaoyou-bilibili/xorm/gen"
)

func main() {
	db, err := driver.NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
	if err != nil {
		panic(err)
	}
	err = gen.GenerateCode(db, []string{"people"}, gen.GenerateConfig{Path: "./query", Pkg: "query"})
	fmt.Println(err)
}

```

### 命令行生成
使用下面的命令安装一下工具
```shell
go get github.com/xiaoyou-bilibili/xorm@latest
go install github.com/xiaoyou-bilibili/xorm/tools/xorm@latest
```
> 注意：这里需要把 GOPATH/bin 目录放到环境变量里去，要不然会找不到命令

我们可以查看相关的帮助信息

```shell
PS D:\code\go\xorm> xorm --help
Usage of C:\APP\go\path\bin\xorm.exe:
  -db string
        数据库名
  -ip string
        数据库ip (default "localhost")
  -passwd string
        数据库密码
  -path string
        生成代码的路径 (default "./query")
  -pkg string
        包名 (default "query")
  -port string
        数据库端口 (default "3306")
  -table string
        表名,多个表用,隔开
  -user string
        数据库用户名 (default "root")
```

比如我们使用下面这样的命令就可以自动在query目录下生成代码了
```shell
xorm -ip 192.168.1.10 -port 8006 -user root -passwd xiaoyou -db demo -table people -path ./query -pkg query
```

## 使用
使用也有两种方式
- 原始调用：这个方法最灵活，不过使用起来也比较麻烦
- 代码生成：代码生成会自动生成model代码以及CRUD代码，使用起来很方便

我们可以新建一个测试表
```sql
CREATE TABLE `people` (
    `id` int NOT NULL,
    `name` varchar(255) NOT NULL,
    `age` int DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

### 原始调用

```go
db, err := driver.NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
if err != nil {
    panic(err)
}

//新增数据
res, err := db.Create("people", map[string]interface{}{
    "id":   1,
    "name": "测试",
    "age":  20,
})
fmt.Println(res, err)

// 删除id为1的数据
res, err = db.Delete("people", []*driver.ConditionInfo{
    {FieldName: "id", FieldValue: []interface{}{1}, Option: driver.ConditionOptionEq},
})

// 把id为1的数据的name修改为测试 2
res,err = db.Update("people", map[string]interface{}{"name": "测试2"}, []*driver.ConditionInfo{{FieldName: "id", FieldValue: []interface{}{1}, Option: driver.ConditionOptionEq}})

// 查找id为1的数据，这里需要定义一下模型，对应的数据库字段为 xorm
type People struct {
    Id   int64   `xorm:"id" json:"id"`
    Name *string `xorm:"name" json:"name"`
    Age  *int64  `xorm:"age" json:"age"`
}
res2, err := db.Find("people", driver.FindInfo{
    Conditions: []*driver.ConditionInfo{{FieldName: "id", FieldValue: []interface{}{1}, Option: driver.ConditionOptionEq}},
}, reflect.TypeOf(&People{}))
for _, people := range res2.([]*People) {
    fmt.Println(utils.Interface2String(people))
}

// 事务操作
err = db.Transaction(func(tx driver.DbInstance) error {
    // 使用方法和上面一样，只是需要把db改成tx，这样才能在事务里操作
    tx.Create("people", map[string]interface{}{})
    // 这个函数返回错误就说明需要回滚，如果返回nil就会提交事务
    return err
})
```

### 代码生成

```go	
db, err := driver.NewMysqlDevice("192.168.1.10", 8006, "root", "xiaoyou", "demo")
if err != nil {
    panic(err)
}

//初始化query
q := query.NewQuery(db)
// 获取query对应的模型
people := q.People

// 新增数据
age := int64(11)
model := &query.People{
    Id:   3,
    Name: "小美",
    Age:  &age,
}
err = people.Create(model)

// 删除id为3的数据
res, err := people.Where(people.Id.Eq(3)).Delete()

// 把id为3的数据的名字改为小美2
res, err = people.Where(people.Id.Eq(3)).Update(people.Name, "小美2")
// 如果要修改多个字段可以这样
res, err = people.Where(people.Id.Eq(3)).UpdateMulti(map[field.IField]interface{}{
    people.Name: "小美2",
    people.Age:  22,
})

// 查询所有id大于1的数据，结果根据年龄降序排序，offset为1，同时只返回一条数据
peoples, err := people.Where(people.Id.Gt(1)).OrderBy(people.Age.Desc()).Limit(1).Offset(1).Find()
for _, people := range peoples {
    fmt.Println(fmt.Println(people))
}

fmt.Println(res, err)
```

## 项目结构&原理

### 项目结构
```shell
├─driver # 驱动相关，如果有新驱动就可以在这个目录里面新建一个文件并实现对于方法即可
├─gen # 代码生成相关，和代码生成有关的逻辑都在这里
│  ├─do # 这里封装了一下驱动的接口，提供更高层次的api
│  └─field # 字段代理，对常用的一些字段进行代理，提供常用判断条件
├─template # 代码模板
├─tools # 工具
│  └─xorm # 命令行代码生成工具
└─utils # 工具类
```

### 多驱动支持

为了实现多数据源，就必须要操作进行抽象，一般数据库我们需要关注的有下面几个

- 返回的字段
- 查询条件
- 排序条件
- limit
- offset

比如我这里就把查询条件的结构体设置如下
```go
type ConditionInfo struct {
	Or         bool            // 是否为或，默认为and
	FieldName  string          // 字段名称
	Option     ConditionOption // 具体操作
	FieldValue []interface{}   // 判断的值
}
```
因为查询条件分为很多种，这里我们可以使用枚举的方式来把这些操作都给列出来
```go
// ConditionOption 具体操作
type ConditionOption uint

const (
    ConditionOptionLike     ConditionOption = iota //  like
    ConditionOptionNLike                           //  not like
    ConditionOptionEq                              // =
    ConditionOptionNeq                             // !=
    ConditionOptionGt                              // >
    ConditionOptionGte                             // >=
    ConditionOptionLt                              // <
    ConditionOptionLte                             // <=
    ConditionOptionIn                              // in
    ConditionOptionNIn                             // not in
    ConditionOptionBetween                         // between
    ConditionOptionNBetween                        // not between
)
```

然后我这里为了方便起见，就直接把增删改查都直接抽象为单个接口，定义如下
```go
type DbInstance interface {
	// DataBaseName 获取数据库名称
	DataBaseName() string
	// SqlType 当前数据库类型
	SqlType() string
	// Create 新增数据
	Create(table string, fields map[string]interface{}) (affected int64, err error)
	// Delete 删除数据
	Delete(table string, conditions []*ConditionInfo) (affected int64, err error)
	// Update 更新数据
	Update(table string, fields map[string]interface{}, conditions []*ConditionInfo) (affected int64, err error)
	// Find 查找数据
	Find(table string, info FindInfo, p reflect.Type) (interface{}, error)
	// Transaction 事务操作
	Transaction(handle func(tx DbInstance) error) error
	// RowQuery 原始查询操作
	RowQuery(sql string, args ...interface{}) (*sql.Rows, error)
	// RowExec 原始执行操作
	RowExec(sql string, args ...interface{}) (int64, error)
}
```

而具体驱动要做的就是把这些定义的结构体翻译为对应数据库的SQL语句，因为我这个考虑的非常简单，所以实现起来也比较简单，实际情况肯定比这个更加复杂

### 结构体映射

本项目的第二个难点就是如何把数据库查询到的数据给映射为go里面对应的结构体，这块也是花了较多心思的地方，具体实现在 `utils/scan.go`里。

这个功能的核心就是反射，通过反射来获取结构体的字段，提取tag，然后tag和sql查询到的字段进行一一对应，详细过程这里就不多说，代码里面写的很详细了

### 代码生成

之前我一直以为go语言代码生成很高级，其实实现意外很简单，那就是模板，因为我们生成的代码结构其实是非常固定的，所以我们只需要提前定义好结构，然后需要改变的地方就可以使用模板语法自动生成。比如我这个项目所有的模板都在`template`里

最后就是把生成的代码写入到具体文件就行了



