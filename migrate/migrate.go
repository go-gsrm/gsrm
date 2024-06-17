package migrate

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/go-gsrm/gsrm/utils"
)

func AutoMigrate[T any](db *sql.DB) {
	db.Query(GenerateTableSQL[T]())
}

func GenerateTableSQL[T any]() string {
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
	return field.Name + " " + GenerateFieldTypeSQLByType(field.Type)
}

func GenerateFieldTypeSQLByType(t reflect.Type) (typeSQL string) {
	if t.Kind() == reflect.Ptr {
		return GenerateFieldTypeSQLByKind(t.Elem().Kind())
	} else {
		return GenerateFieldTypeSQLByKind(t.Kind()) + " NOT NULL"
	}
}

func GenerateFieldTypeSQLByKind(k reflect.Kind) string {
	switch k {
	case reflect.String:
		return "varchar(255)"
	case reflect.Int:
		return "bigint"
	case reflect.Int32:
		return "bigint"
	case reflect.Int64:
		return "bigint"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.Uint:
		return "int unsigned"
	case reflect.Float64:
		return "DOUBLE"
	case reflect.Float32:
		return "DOUBLE"
	default:
		panic("unsupported type")
	}
}
