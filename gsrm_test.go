package gsrm

import (
	"testing"

	"github.com/go-gsrm/gsrm/migrate"
)

var db *DB

func init() {

}

func TestOpen(t *testing.T) {
	db = Open("mysql", "jarvis:jarvis@/jarvis?charset=utf8")
}

type testInsertStruct struct {
	ID   int
	Name string
	Rate float64
	Age  *int
}

func TestInsert(t *testing.T) {
	migrate.AutoMigrate[testInsertStruct](db.DB)
	age := 123
	Insert(db.DB, testInsertStruct{
		ID:   1,
		Name: "test",
		Rate: 0.1,
		Age:  &age,
	})
	Insert(db.DB, testInsertStruct{
		ID:   1,
		Name: "test",
		Rate: 32123,
		Age:  nil,
	})
}

type testStruct struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Price  int    `json:"price" gorm:"type:int"`
	Name   string `json:"name" gorm:"type:varchar(255)"`
	Age    uint   `json:"age" gorm:"type:int"`
	Height *uint  `json:"height" gorm:"type:int"`
}

func TestCreateTable(t *testing.T) {
	// CreateTable()
	migrate.AutoMigrate[testStruct](db.DB)
}
