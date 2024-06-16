package utils

import "testing"

func TestGetFieldsNameByInstance(t *testing.T) {
	type TestGetFieldsNameStruct struct {
		ID   int
		Name string
	}
	var test TestGetFieldsNameStruct = TestGetFieldsNameStruct{
		ID:   1,
		Name: "test",
	}
	t.Log(GetFieldsNameByInstance(test))
}
