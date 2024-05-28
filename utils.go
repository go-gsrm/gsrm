package gsrm

import (
	"reflect"
	"strings"
)

func generateInsertSQL[T any](t T) string {
	var fieldsName []string = make([]string, 0)
	typeOfT := reflect.TypeOf(t)
	for i := 0; i < typeOfT.NumField(); i++ {
		field := typeOfT.Field(i)
		fieldsName = append(fieldsName, field.Name)
	}
	fieldsNameString := strings.Join(fieldsName, ",")
	valueOf := reflect.ValueOf(t)
	tableName := valueOf.Type().Name()
	sql := "INSERT INTO " + tableName + " (" + fieldsNameString + ") VALUES (" + ");"
	return sql
}
