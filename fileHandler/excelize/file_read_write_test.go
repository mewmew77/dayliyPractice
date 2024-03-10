package excelize

import "testing"

func TestWrite(t *testing.T) {
	fileName := "test.xlsx"
	data := [][]string{
		{"列1", "列2"},
		{"val1", "val2"},
	}
	sheetName := "Sheet1"
	if err := Write(fileName, sheetName, data); err != nil {
		t.Fatal("write file failed, err = ", err)
	}
}

func TestRead(t *testing.T) {
	fileName := "test.xlsx"
	sheetName := "Sheet1"
	data, err := Read(fileName, sheetName)
	if err != nil {
		t.Fatal("read file failed, err = ", err)
	}
	t.Log("success : ", data)
}
