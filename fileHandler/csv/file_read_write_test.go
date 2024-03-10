package csv

import "testing"

func TestWrite(t *testing.T) {
	fileName := "test.csv"
	data := [][]string{
		{"列1", "列2"},
		{"val1", "val2"},
	}
	if err := Write(fileName, data); err != nil {
		t.Fatal("write file failed, err = ", err)
	}
}

func TestRead(t *testing.T) {
	fileName := "test.csv"
	data, err := Read(fileName)
	if err != nil {
		t.Fatal("read file failed, err = ", err)
	}
	t.Log("success : ", data)
}
