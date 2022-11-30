package do

import "github.com/xiaoyou-bilibili/xorm/driver"

type Do struct {
	table      string
	columns    []string
	conditions []*driver.ConditionInfo
	order      []*driver.OrderInfo
	limit      *int
	offset     *int
	fields     map[string]interface{}
}

func (do Do) _StartSelect() {

}
