package gsrm

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/go-gsrm/gsrm/utils"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func GenericDB[T any](db *sql.DB) *GDB[T] {
	return &GDB[T]{db}
}

type GDB[T any] struct {
	*sql.DB
}

func (db *GDB[T]) Query(query string, args ...any) {
	db.DB.Query(query, args...)
}

func (db *GDB[T]) Exec(query string, args ...any) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *GDB[T]) Insert(data T) (T, error) {
	return Insert(db.DB, data)
}

func (db *GDB[T]) InsertMany(data ...T) []T {
	return InsertMany(db.DB, data...)
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

func Insert[T any](db *sql.DB, data T) (T, error) {
	tableName := utils.GetTableNameByInstance(data)
	fieldsName := utils.GetFieldsNameByInstance(data)
	valueOf := reflect.ValueOf(data)
	method := valueOf.MethodByName("GsrmBeforeInsert")
	if method.IsValid() {
		// TODO: error handling
		results := method.Call(nil)
		newData := results[0].Interface()
		data = newData.(T)
	}
	valueOf = reflect.ValueOf(data)
	query := "INSERT INTO " + tableName + " ("
	query += strings.Join(fieldsName, ",")
	placeholder := strings.Repeat("?,", len(fieldsName))
	query += ") VALUES (" + placeholder[:len(placeholder)-1]
	query += ")"
	args := make([]any, valueOf.NumField())
	for i := 0; i < valueOf.NumField(); i++ {
		args[i] = valueOf.Field(i).Interface()
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return data, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return data, err
	}
	return data, err
}

func InsertMany[T any](db *sql.DB, data ...T) []T {

	db.Query("INSERT INTO fast_chat_contexts")
	return data
}

func First[T any](db *sql.DB) T {
	var t T
	db.Query("SELECT * FROM fast_chat_contexts")
	return t
}

func List[T any](db *sql.DB) []T {
	var t []T
	db.Query("SELECT * FROM fast_chat_contexts")
	return t
}

func Delete[T any](db *sql.DB, t ...T) int64 {
	db.Query("DELETE FROM fast_chat_contexts")
	return 0
}

func Update[T any](db *sql.DB, t T) T {
	db.Query("UPDATE fast_chat_contexts")
	return t
}

func UpdateWithMap[T any](db *sql.DB, t T, m map[string]any) T {
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
