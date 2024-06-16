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

func TestInsert(t *testing.T) {
	Insert(db, testStruct{
		ID: 1,
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
