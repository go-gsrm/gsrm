package migrate

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-gsrm/gsrm/utils"
	"github.com/huandu/go-sqlbuilder"
)

func AutoMigrate[T any](db *sql.DB) error {
	_, err := db.Exec(GenerateTableSQL[T]())
	return err
}

func GenerateTableSQL[T any]() string {
	var structType T
	builder := sqlbuilder.CreateTable(utils.GetTableNameByInstance(structType)).
		IfNotExists()
	typeOf := reflect.TypeOf(structType)
	for _, fieldProperty := range utils.GetFieldPropertiesByType(typeOf) {
		builder.Define(fieldProperty.Define()...)
	}
	if primaryKeys := utils.GetPrimaryKeyColumnsByType(typeOf); len(primaryKeys) > 0 {
		builder.Define("PRIMARY KEY", "("+strings.Join(primaryKeys, ",")+")")
	}
	sql, _ := builder.Build()
	fmt.Println(sql)
	return sql
}
