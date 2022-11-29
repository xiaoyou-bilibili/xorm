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
	for index, condition := range info {
		if len(condition.FieldValue) == 0 {
			continue
		}
		if index != 1 {
			if condition.Or {
				conditions.WriteString(" AND ")
			} else {
				conditions.WriteString(" OR ")
			}
		}
		switch condition.Option {
		case ConditionOptionLike:
			conditions.WriteString(condition.FieldName + " LIKE ?")
			values = append(values, condition.FieldValue[0])
		case ConditionOptionNLike:

		}
	}
	return conditions.String()
}

func (d *MysqlDriver) Create(table string, fields map[string]interface{}) (affected int64, err error) {
	field, placeholder, value := d.fieldsConcat(fields)
	rowSql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, field, placeholder)
	log.Printf("sql is %s", rowSql)
	res, err := d.db.Exec(rowSql, value...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *MysqlDriver) Delete(table string, conditions []*ConditionInfo) (affected int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *MysqlDriver) Update(table string, fields map[string]interface{}, conditions []*ConditionInfo) (affected int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *MysqlDriver) Find(table string, info FindInfo, res interface{}) error {
	//TODO implement me
	panic("implement me")
}
