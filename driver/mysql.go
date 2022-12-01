package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaoyou-bilibili/xorm/utils"
	"log"
	"reflect"
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
	tx *sql.Tx
}

func (d *MysqlDriver) insertFieldsConcat(fields map[string]interface{}) (string, string, []interface{}) {
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

func (d *MysqlDriver) updateFieldConcat(fields map[string]interface{}) (string, []interface{}) {
	field := make([]string, 0, len(fields))
	value := make([]interface{}, 0, len(fields))
	for k, v := range fields {
		field = append(field, fmt.Sprintf("`%s`=?", k))
		value = append(value, v)
	}
	return strings.Join(field, ","), value
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
		conditions.WriteString(fmt.Sprintf("`%s` %s %s", condition.FieldName, option, placeholder))
		values = append(values, condition.FieldValue...)
	}
	return conditions.String(), values
}

func (d *MysqlDriver) orderConcat(info []*OrderInfo) string {
	var orders []string
	for _, order := range info {
		asc := "ASC"
		if order.Desc {
			asc = "DESC"
		}
		orders = append(orders, fmt.Sprintf("%s %s", order.FieldName, asc))
	}
	return strings.Join(orders, ",")
}

func (d *MysqlDriver) rowsProcess(result sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (d *MysqlDriver) Create(table string, fields map[string]interface{}) (affected int64, err error) {
	field, placeholder, values := d.insertFieldsConcat(fields)
	rowSql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, field, placeholder)
	log.Printf("sql is %s", rowSql)
	if d.tx != nil {
		return d.rowsProcess(d.tx.Exec(rowSql, values...))
	}
	return d.rowsProcess(d.db.Exec(rowSql, values...))
}

func (d *MysqlDriver) Delete(table string, conditions []*ConditionInfo) (affected int64, err error) {
	condition, values := d.conditionConcat(conditions)
	rowSql := "DELETE FROM " + table
	if len(conditions) > 0 {
		rowSql += " WHERE " + condition
	}
	log.Printf("sql is %s", rowSql)
	if d.tx != nil {
		return d.rowsProcess(d.tx.Exec(rowSql, values...))
	}
	return d.rowsProcess(d.db.Exec(rowSql, values...))
}

func (d *MysqlDriver) Update(table string, fields map[string]interface{}, conditions []*ConditionInfo) (affected int64, err error) {
	field, values := d.updateFieldConcat(fields)
	rowSql := fmt.Sprintf("UPDATE %s SET %s", table, field)
	if len(conditions) > 0 {
		condition, values2 := d.conditionConcat(conditions)
		rowSql += " WHERE " + condition
		values = append(values, values2...)
	}
	log.Printf("sql is %s", rowSql)
	if d.tx != nil {
		return d.rowsProcess(d.tx.Exec(rowSql, values...))
	}
	return d.rowsProcess(d.db.Exec(rowSql, values...))
}

func (d *MysqlDriver) Find(table string, info FindInfo, modelType reflect.Type) (interface{}, error) {
	var values []interface{}
	rowSql := strings.Builder{}
	field := []string{"*"}
	if len(info.Columns) > 0 {
		field = []string{}
		for _, column := range info.Columns {
			field = append(field, fmt.Sprintf("`%s`", column))
		}
	}
	rowSql.WriteString(fmt.Sprintf("SELECT %s FROM %s", strings.Join(field, ","), table))
	if len(info.Conditions) > 0 {
		condition, values2 := d.conditionConcat(info.Conditions)
		rowSql.WriteString(" WHERE " + condition)
		values = append(values, values2...)
	}
	if len(info.Orders) > 0 {
		condition := d.orderConcat(info.Orders)
		rowSql.WriteString(" ORDER BY " + condition)
	}
	if info.Limit != nil {
		rowSql.WriteString(fmt.Sprintf(" LIMIT %d", *info.Limit))
	}
	if info.Offset != nil {
		rowSql.WriteString(fmt.Sprintf(" OFFSET %d", *info.Offset))
	}
	log.Printf("sql is %s", rowSql.String())
	rows, err := d.db.Query(rowSql.String(), values...)
	if err != nil {
		return nil, err
	}
	return utils.ConvertRows2Struct(rows, modelType)
}

func (d *MysqlDriver) Transaction(handle func(tx DbInstance) error) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	instance := &MysqlDriver{tx: tx}
	err = handle(instance)
	if err != nil {
		log.Printf("handle err %v, rollback", err)
		return tx.Rollback()
	}
	return tx.Commit()
}
