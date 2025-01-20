package companyagencies

import (
	"encoding/csv"
	"os"
)

func connectToCSV() (*os.File, error) {
	csv, err := os.Open("data/data.csv")
	if err != nil {
		return nil, err
	}
	defer csv.Close()

	return csv, nil
}

func readCSVFile(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil
}
