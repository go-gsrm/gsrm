package gsrm

import (
	"database/sql"
	"strings"

	"github.com/go-gsrm/gsrm/utils"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

type Column struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func Open(driverName, dataSourceName string) *DB {
	b, _ := sql.Open(driverName, dataSourceName)
	return &DB{b}
}

func Insert[T any](db *DB, t T) T {
	var structType T
	tableName := utils.GetTableNameByInstance(structType)
	fieldsName := utils.GetFieldsNameByInstance(structType)
	query := "INSERT INTO " + tableName + " ("
	query += strings.Join(fieldsName, ",")
	placeholder := strings.Repeat("?,", len(fieldsName))
	query += ") VALUES (" + placeholder[:len(placeholder)-1]
	query += ")"

	// db.Query(query)
	return t
}

func InsertMany[T any](db DB, t ...T) []T {
	db.Query("INSERT INTO fast_chat_contexts")
	return t
}

func First[T any](db DB) T {
	var t T
	db.Query("SELECT * FROM fast_chat_contexts")
	return t
}

func List[T any](db DB) []T {
	var t []T
	db.Query("SELECT * FROM fast_chat_contexts")
	return t
}

func Delete[T any](db DB, t ...T) int64 {
	db.Query("DELETE FROM fast_chat_contexts")
	return 0
}

func Update[T any](db DB, t T) T {
	db.Query("UPDATE fast_chat_contexts")
	return t
}

func UpdateWithMap[T any](db DB, t T, m map[string]any) T {
	db.Query("UPDATE fast_chat_contexts")
	return t
}

type Repostitory[T any] struct {
	DB DB
}

func (r Repostitory[T]) Insert(t T) T {
	r.DB.Query("INSERT INTO fast_chat_contexts")
	return t
}
