package migrate

import (
	"testing"
)

type TestStruct struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	Age  uint   `json:"age" gorm:"type:int"`
}

func (TestStruct) TableName() string {
	return "TestStruct_test_table"
}

func TestGenerateTableSQL(t *testing.T) {
	// TODO
	sql := GenerateTableSQL[TestStruct]()
	t.Log(sql)
}
