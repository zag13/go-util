package goutils

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name string
	Date string
}

var testData = []TestStruct{
	{Name: "test1", Date: "01/01/2001"},
	{Name: "test2", Date: "02/02/2002"},
}

func TestJsonDump(t *testing.T) {
	n, err := JsonDump("./test.json", testData)
	if err != nil {
		t.Errorf("JsonDump() failed: %s", err)
	}
	if n == 0 {
		t.Errorf("JsonDump() failed: no data written")
	}
}

func TestJsonLoad(t *testing.T) {
	var data []TestStruct
	err := JsonLoad("./test.json", &data)
	if err != nil {
		t.Errorf("JsonLoad() failed: %s", err)
	}
	if reflect.DeepEqual(data, testData) == false {
		t.Errorf("JsonLoad() failed: data mismatch")
	}
}
