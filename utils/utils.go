package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func GetTableNameByInstance[T any](t T) string {
	valueOf := reflect.ValueOf(t)
	tableName := valueOf.Type().Name()
	method := valueOf.MethodByName("TableName")
	if method.IsValid() {
		values := method.Call(nil)
		if len(values) > 0 {
			tableName = values[0].String()
		}
	}
	return strings.ToLower(tableName)
}

func GetFieldsNameByInstance[T any](t T) []string {
	return GetFieldsNameByType(reflect.TypeOf(t))
}

func GetFieldsNameByType(t reflect.Type) []string {
	fieldsName := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fieldsName[i] = GetFieldName(t.Field(i))
	}
	return fieldsName
}

func GetFieldPropertiesByType(t reflect.Type) []FieldProperty {
	fieldProperties := make([]FieldProperty, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fieldProperties[i] = GetFieldProperty(t.Field(i))
	}
	return fieldProperties
}

func GetFieldProperty(t reflect.StructField) FieldProperty {
	return FieldProperty{
		Name: GetFieldName(t),
		Type: GenerateFieldTypeSQLByKind(t.Type.Kind()),
	}
}

func GetFieldName(t reflect.StructField) string {
	return t.Name
}

func GetTagValue(t reflect.StructField, key string) string {
	tag := GetTag(t)
	fmt.Println(tag)
	reg, _ := regexp.Compile(".*" + key + ":(.*);")
	fmt.Println(reg.FindStringSubmatch(tag))
	return ""
	// return reg.FindStringSubmatch(tag)[1]
}

func GetTagType(t reflect.StructField) string {
	return GetTagValue(t, "type")
}

func GetTag(t reflect.StructField) string {
	return t.Tag.Get("gsrm")
}

func GetPrimaryKeyColumnsByType(t reflect.Type) []string {
	var primaryKey []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if IsPromaryKey(field) {
			primaryKey = append(primaryKey, GetFieldName(field))
		}
	}
	return primaryKey
}

func IsPromaryKey(field reflect.StructField) bool {
	reg, _ := regexp.Compile(".*primaryKey.*")
	tag := GetTag(field)
	return reg.Match([]byte(tag))
}

func GetPlaceholder[T any]() string {
	var structType T
	return GetPlaceholderByInstance(structType)
}

func GetPlaceholderByInstance[T any](t T) string {
	placeholder := strings.Repeat("?,", len(GetFieldsNameByInstance(t)))
	return "(" + placeholder[:len(placeholder)-1] + ")"
}

func GetArgsByValueOf(valueOf reflect.Value) []any {
	args := make([]any, valueOf.NumField())
	for i := 0; i < valueOf.NumField(); i++ {
		args[i] = valueOf.Field(i).Interface()
	}
	return args
}

func ExecBeforeInsert[T any](data T) (T, error) {
	valueOf := reflect.ValueOf(data)
	method := valueOf.MethodByName("GsrmBeforeInsert")
	if method.IsValid() {
		// TODO: error handling
		results := method.Call(nil)
		newData := results[0].Interface()
		data = newData.(T)
	}
	return data, nil
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
