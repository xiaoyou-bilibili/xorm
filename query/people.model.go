package query

type People struct {
	Id   int64  `xorm:"id"`
	Name string `xorm:"name"`
	Age  int64  `xorm:"age"`
}

func (p *People) TableName() string {
	return "people"
}
