package utils

import "reflect"

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

func GetFieldsNameByInstance[T any](t T) []string {
	typeOf := reflect.TypeOf(t)
	fieldsName := make([]string, typeOf.NumField())
	for i := 0; i < typeOf.NumField(); i++ {
		fieldsName[i] = typeOf.Field(i).Name
	}
	return fieldsName
}
