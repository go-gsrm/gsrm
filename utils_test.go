package gsrm

import "testing"

type TestGenerateInertSQLStruct struct {
	ID   int
	Name string
}

func TestGenerateInertSQL(t *testing.T) {
	var test TestGenerateInertSQLStruct = TestGenerateInertSQLStruct{
		ID:   1,
		Name: "test",
	}
	t.Log(generateInsertSQL(test))
}
