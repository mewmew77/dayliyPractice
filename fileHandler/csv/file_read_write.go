package csv

import (
	"encoding/csv"
	"os"
)

func Write(fileName string, content [][]string) error {
	newFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer newFile.Close()
	_, err = newFile.WriteString("\xEF\xBB\xBF")
	if err != nil {
		return err
	}
	writer := csv.NewWriter(newFile)
	for _, line := range content {
		if err := writer.Write(line); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func Read(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return content, nil
}
