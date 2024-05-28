package migrate

import (
	"fmt"
	"reflect"
	"strings"
)

func IsValid[T any]() bool {
	var target T
	fmt.Println(target)
	return true
}

func GenerateTableSQL[T any]() string {
	if !IsValid[T]() {
		panic("table Invalid type")
	}
	var structType T
	sql := "CREATE TABLE IF NOT EXISTS "
	valueOf := reflect.ValueOf(structType)
	tableName := valueOf.Type().Name()
	method := valueOf.MethodByName("TableName")
	if method.IsValid() {
		values := method.Call(nil)
		if len(values) > 0 {
			tableName = values[0].String()
		}
	}
	typeOf := reflect.TypeOf(structType)
	sql += tableName + " ("
	fieldSQL := make([]string, 0)
	for i := 0; i < typeOf.NumField(); i++ {
		fieldSQL = append(fieldSQL, GenerateFieldSQL(typeOf.Field(i)))
	}
	sql += strings.Join(fieldSQL, ",")
	sql += ") ENGINE InnoDB;"
	return sql
}

func GenerateFieldSQL(field reflect.StructField) string {
	typeSQL := ""
	// unsigned := false
	switch field.Type.Kind() {
	case reflect.String:
		typeSQL = "varchar(255)"
	case reflect.Int:
		typeSQL = "bigint"
	case reflect.Int32:
		typeSQL = "bigint"
	case reflect.Int64:
		typeSQL = "bigint"
	case reflect.Bool:
		typeSQL = "BOOLEAN"
	case reflect.Uint:
		typeSQL = "int unsigned"
	case reflect.Float64:
		typeSQL = "DOUBLE"
	case reflect.Float32:
		typeSQL = "DOUBLE"
	default:
		panic("unsupported type")
	}
	return field.Name + " " + typeSQL
}
