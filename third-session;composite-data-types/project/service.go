package companyagencies

import (
	"os"
	"strconv"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
)

func list(csvFile *os.File, region string) ([]Agency, error) {
	rows, err := readCSVFile(csvFile)
	if err != nil {
		return nil, err
	}

	var agencies []Agency
	for _, row := range rows {
		agency, err := parseAgencyFromCSV(row)
		if err != nil {
			return nil, err
		}

		if region == row[1] {
			agencies = append(agencies, *agency)
		}

	}

	return agencies, nil
}

func get(csvFile *os.File, id int64) (*Agency, error) {
	stringID := strconv.Itoa(int(id))
	var agency *Agency

	rows, err := readCSVFile(csvFile)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		if row[0] == stringID {
			agency, err = parseAgencyFromCSV(row)
			if err != nil {
				return nil, err
			}
		}
	}

	return agency, nil
}

func parseAgencyFromCSV(CSVRow []string) (*Agency, error) {
	id, err := strconv.Atoi(CSVRow[0])
	if err != nil {
		return nil, err
	}

	time, err := time.Parse(TimeLayout, CSVRow[5])

	employeesCount, err := strconv.Atoi(CSVRow[6])
	if err != nil {
		return nil, err
	}

	return &Agency{
		ID:                   int64(id),
		Region:               CSVRow[1],
		Name:                 CSVRow[2],
		Address:              CSVRow[3],
		PhoneNumber:          CSVRow[4],
		RegistrationDate:     time,
		AgencyEmployeesCount: int32(employeesCount),
	}, nil
}

func create(csvFile *os.File, agency *Agency) (int64, error) {
	rows, err := readCSVFile(csvFile)
	if err != nil {
		return 0, err
	}

	agency.ID = int64(len(rows) + 1)
	agency.RegistrationDate = time.Now()
	convertedAgency := convertAgencyToString(agency)

	err = writeAgencyToCSV(csvFile, convertedAgency)
	if err != nil {
		return 0, err
	}

	return agency.ID, nil
}

func convertAgencyToString(agency *Agency) []string {
	stringID := strconv.Itoa(int(agency.ID))
	stringEmployeesCount := strconv.Itoa(int(agency.AgencyEmployeesCount))

	return []string{stringID, agency.Region, agency.Name, agency.Address, agency.RegistrationDate.String(), stringEmployeesCount}
}
