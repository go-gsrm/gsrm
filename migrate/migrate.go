package migrate

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-gsrm/gsrm/utils"
)

func AutoMigrate[T any](db *sql.DB) {
	db.Query(GenerateTableSQL[T]())
}

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
	tableName := utils.GetTableNameByInstance(structType)
	typeOf := reflect.TypeOf(structType)
	sql += tableName + " ("
	fieldSQL := make([]string, typeOf.NumField())
	for i := 0; i < typeOf.NumField(); i++ {
		fieldSQL[i] = GenerateFieldSQL(typeOf.Field(i))
	}
	sql += strings.Join(fieldSQL, ",")
	sql += ") ENGINE InnoDB;"
	return sql
}

func GenerateFieldSQL(field reflect.StructField) string {
	return field.Name + " " + GenerateFieldTypeSQLByKind(field.Type)
}

func GenerateFieldTypeSQLByKind(t reflect.Type) (typeSQL string) {
	notNull := true
	switch t.Kind() {
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
	case reflect.Pointer:
		notNull = false
		typeSQL = GenerateFieldTypeSQLByKind(t.Elem())
	default:
		panic("unsupported type")
	}
	if !notNull {
		typeSQL += " NOT NULL"
	}
	return typeSQL
}
