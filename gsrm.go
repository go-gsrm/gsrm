package gsrm

import (
	"database/sql"
	"reflect"

	"github.com/go-gsrm/gsrm/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/huandu/go-sqlbuilder"
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

func (db *GDB[T]) InsertMany(data ...T) ([]T, error) {
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
	data, _ = utils.ExecBeforeInsert(data)
	builder := generateSQLBuilderByInstence(data)
	builder.Values(generateValuesByInstence(data)...)
	sql, args := builder.Build()
	_, err := db.Exec(sql, args...)
	if err != nil {
		return data, err
	}
	return data, err
}

func InsertMany[T any](db *sql.DB, dataList ...T) ([]T, error) {
	var err error
	if len(dataList) == 0 {
		return dataList, nil
	}
	builder := generateSQLBuilderByInstence(dataList[0])
	output := make([]T, len(dataList))
	for i, data := range dataList {
		data, err = utils.ExecBeforeInsert(data)
		if err != nil {
			return dataList, err
		}
		builder.Values(generateValuesByInstence(data)...)
		output[i] = data
	}
	sql, args := builder.Build()
	_, err = db.Exec(sql, args...)
	return output, err
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

func generateSQLBuilderByInstence[T any](data T) *sqlbuilder.InsertBuilder {
	builder := sqlbuilder.NewInsertBuilder()
	builder.InsertInto(utils.GetTableNameByInstance(data))
	typeOf := reflect.TypeOf(data)
	var cloumnsName []string = make([]string, 0)
	for i := 0; i < typeOf.NumField(); i++ {
		cloumnsName = append(cloumnsName, typeOf.Field(i).Name)
	}
	builder.Cols(cloumnsName...)
	return builder
}

func generateValuesByInstence[T any](data T) []interface{} {
	valueOf := reflect.ValueOf(data)
	var values []interface{} = make([]interface{}, 0)
	for i := 0; i < valueOf.NumField(); i++ {
		values = append(values, valueOf.Field(i).Interface())
	}
	return values
}
