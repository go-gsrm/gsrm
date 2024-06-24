package utils

import (
	"reflect"
	"regexp"
	"strings"
)

func GetTableName[T any]() string {
	var structType T
	return GetTableNameByInstance(structType)
}

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
	return tableName
}

func GetFieldsName[T any]() []string {
	var structType T
	return GetFieldsNameByInstance(structType)
}

func GetFieldsNameByInstance[T any](t T) []string {
	return GetFieldsNameByType(reflect.TypeOf(t))
}

func GetFieldsNameByType(t reflect.Type) []string {
	fieldsName := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		fieldsName[i] = t.Field(i).Name
	}
	return fieldsName
}

func GetTag(t reflect.StructField) string {
	return t.Tag.Get("gsrm")
}

func GetPrimaryKeyColumnsByType(t reflect.Type) []string {
	var primaryKey []string
	reg, _ := regexp.Compile(".*primaryKey.*")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := GetTag(field)
		if reg.Match([]byte(tag)) {
			primaryKey = append(primaryKey, field.Name)
		}
	}
	return primaryKey
}

func GetPlaceholder[T any]() string {
	var structType T
	return GetPlaceholderByInstance(structType)
}

func GetPlaceholderByInstance[T any](t T) string {
	fieldsName := GetFieldsNameByInstance(t)
	placeholder := strings.Repeat("?,", len(fieldsName))
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
