package companyagencies

import (
	"encoding/csv"
	"io"
	"os"
)

func ConnectToCSV() (*os.File, error) {
	csv, err := os.OpenFile("data/data.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return csv, nil
}

func readCSVFile(file *os.File) ([][]string, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func writeAgencyToCSV(csvFile *os.File, agency []string) error {
	_, err := csvFile.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	err = csvwriter.Write(agency)
	if err != nil {
		return err
	}

	_, err = csvFile.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}
