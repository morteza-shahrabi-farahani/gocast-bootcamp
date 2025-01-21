package companyagencies

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ConnectToCSV() (*os.File, error) {
	csv, err := os.Open("data/data.csv")
	if err != nil {
		return nil, err
	}

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

func writeAgencyToCSV(csvFile *os.File, agency []string) error {
	csvwriter := csv.NewWriter(csvFile)

	fmt.Println(agency)
	err := csvwriter.Write(agency)
	if err != nil {
		return err
	}

	return nil
}
