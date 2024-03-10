package excelize

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func Write(fileName string, sheetName string, content [][]string) error {
	file := excelize.NewFile()
	file.NewSheet(sheetName)
	for i := range content {
		file.SetSheetRow(sheetName, fmt.Sprintf("A%d", i+1), &content[i])
	}
	err := file.SaveAs(fileName)
	if err != nil {
		return err
	}
	return nil
}

func Read(filename string, sheetName string) ([][]string, error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	rows := file.GetRows(sheetName)
	return rows, nil
}
