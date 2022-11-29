package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

// NewMysqlDevice 初始化mysql驱动
func NewMysqlDevice(ip string, port int, username, password, database string) (DbInstance, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, database))
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MysqlDriver{
		db: db,
	}, nil
}

type MysqlDriver struct {
	db *sql.DB
}

func (d *MysqlDriver) fieldsConcat(fields map[string]interface{}) (string, string, []interface{}) {
	field := make([]string, 0, len(fields))
	value := make([]interface{}, 0, len(fields))
	placeholder := make([]string, 0, len(field))
	for k, v := range fields {
		field = append(field, fmt.Sprintf("`%s`", k))
		value = append(value, v)
		placeholder = append(placeholder, "?")
	}
	return strings.Join(field, ","), strings.Join(placeholder, ","), value
}

func (d *MysqlDriver) conditionConcat(info []*ConditionInfo) (string, []interface{}) {
	conditions := strings.Builder{}
	var values []interface{}
	// 操作符和占位符
	option := ""
	placeholder := "?"
	// 遍历所有条件
	for index, condition := range info {
		if len(condition.FieldValue) == 0 {
			continue
		}
		if index != 0 {
			if condition.Or {
				conditions.WriteString(" OR ")
			} else {
				conditions.WriteString(" AND ")
			}
		}
		switch condition.Option {
		case ConditionOptionLike:
			option = "LIKE"
		case ConditionOptionNLike:
			option = "NOT LIKE"
		case ConditionOptionEq:
			option = "="
		case ConditionOptionNeq:
			option = "!="
		case ConditionOptionGt:
			option = ">"
		case ConditionOptionGte:
			option = ">="
		case ConditionOptionLt:
			option = "<"
		case ConditionOptionLte:
			option = "<="
		case ConditionOptionIn, ConditionOptionNIn:
			option = "IN"
			if condition.Option == ConditionOptionNIn {
				option = "NOT " + option
			}
			pl := make([]string, 0, len(condition.FieldValue))
			for range condition.FieldValue {
				pl = append(pl, "?")
			}
			placeholder = fmt.Sprintf("(%s)", strings.Join(pl, ","))
		case ConditionOptionBetween, ConditionOptionNBetween:
			// between条件需要有两个
			if len(condition.FieldValue) != 2 {
				continue
			}
			option = "BETWEEN"
			if condition.Option == ConditionOptionNIn {
				option = "NOT " + option
			}
			placeholder = "? AND ?"
		}
		conditions.WriteString(fmt.Sprintf("%s %s %s", condition.FieldName, option, placeholder))
		values = append(values, condition.FieldValue...)
	}
	return conditions.String(), values
}

func (d *MysqlDriver) rowsProcess(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (d *MysqlDriver) Create(table string, fields map[string]interface{}) (affected int64, err error) {
	field, placeholder, values := d.fieldsConcat(fields)
	rowSql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, field, placeholder)
	log.Printf("sql is %s", rowSql)
	return d.rowsProcess(d.db.Exec(rowSql, values...))
}

func (d *MysqlDriver) Delete(table string, conditions []*ConditionInfo) (affected int64, err error) {
	condition, values := d.conditionConcat(conditions)
	rowSql := fmt.Sprintf("DELETE FROM %s WHERE %s", table, condition)
	log.Printf("sql is %s", rowSql)
	return d.rowsProcess(d.db.Exec(rowSql, values...))
}

func (d *MysqlDriver) Update(table string, fields map[string]interface{}, conditions []*ConditionInfo) (affected int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *MysqlDriver) Find(table string, info FindInfo, res interface{}) error {
	//TODO implement me
	panic("implement me")
}
