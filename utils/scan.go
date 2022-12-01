package utils

import (
	"database/sql"
	"reflect"
	"strings"
)

// ConvertRows2Struct sql row对象转换为结构体
func ConvertRows2Struct(rows *sql.Rows, modelType reflect.Type) (interface{}, error) {
	var scanRes []map[string]interface{}
	// 先把所有的值给取出来
	for rows.Next() {
		columns, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}
		fields := make([]interface{}, 0, len(columns))
		// 拿到所有字段，每个字段对于一个接口类型
		row := map[string]interface{}{}
		for _, column := range columns {
			// 根据字段类型新建一个对于字段
			field := reflect.New(column.ScanType()).Interface()
			row[column.Name()] = field
			fields = append(fields, field)
		}
		// 提取出字段的值
		err = rows.Scan(fields...)
		if err != nil {
			return nil, err
		}
		scanRes = append(scanRes, row)
	}
	length := len(scanRes)
	// 根据传入的elem类型，创建一个新的切片，大小为sql返回的结果
	newSlice := reflect.MakeSlice(reflect.SliceOf(modelType), length, length)
	for index, values := range scanRes {
		// 获取当前切片的值
		item := newSlice.Index(index)
		// 如果传入的是指针类型，那么还需要获取指针指向的值
		if item.Type().Kind() == reflect.Ptr {
			// 先对item进行初始化
			item.Set(reflect.New(item.Type().Elem()))
			// 获取指针指向的值
			item = item.Elem()
		}
		// 解析结构体
		for i := 0; i < item.NumField(); i++ {
			// 获取结构体对应字段的类型和值
			field := item.Type().Field(i)
			filedVal := item.Field(i)
			// 从map映射中找出该字段对应的值
			if value, ok := values[field.Tag.Get("xorm")]; ok {
				// 获取sql扫描出的值
				elem := reflect.ValueOf(value).Elem()
				//fmt.Println("sql type", elem.Type())
				//fmt.Println("field type", field.Type)
				// 根据不同字段类型进行转换
				switch field.Type.Kind() {
				case reflect.Int64:
					filedVal.SetInt(sqlType2Int64(elem))
				case reflect.String:
					filedVal.SetString(sqlType2String(elem))
				}
			}
		}
	}

	return newSlice.Interface(), nil
}

// 下面需要把sql类型直接转换为自己固定的几个类型
func sqlType2Int64(data reflect.Value) int64 {
	switch data.Kind() {
	case reflect.Int32, reflect.Int64: // 如果是自带的类型直接转换即可
		return data.Int()
	case reflect.TypeOf(sql.NullInt64{}).Kind(): // 如果是sql类型，这里需要断言转换
		return data.Interface().(sql.NullInt64).Int64
	}
	return 0
}

func sqlType2String(data reflect.Value) string {
	switch data.Kind() {
	case reflect.String:
		return data.String()
	case reflect.TypeOf(sql.RawBytes{}).Kind():
		return strings.TrimSpace(string(data.Bytes()))
	}
	return ""
}
