package gsrm

import (
	"database/sql"
	"fmt"

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

func Open() *DB {
	b, _ := sql.Open("mysql", "jarvis:jarvis@/jarvis?charset=utf8")
	// b.Exec("CREATE TABLE IF NOT EXISTS hello_wolr (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	// result, _ := b.Exec("Show Columns from fast_chat_contexts")
	rows, _ := b.Query("Show Columns from fast_chat_contexts")
	for rows.Next() {
		var column Column
		rows.Scan(&column.Field, &column.Type, &column.Null, &column.Key, &column.Default, &column.Extra)
		fmt.Println(column)
	}
	return &DB{}
}

func Insert[T any](db DB, t T) T {
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

func Delete[T any](db DB, t T) int64 {
	db.Query("DELETE FROM fast_chat_contexts")
	return 0
}

func Update[T any](db DB, t T) T {
	db.Query("UPDATE fast_chat_contexts")
	return t
}
