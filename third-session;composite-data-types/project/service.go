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
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}

		time, err := time.Parse(TimeLayout, row[5])

		employeesCount, err := strconv.Atoi(row[6])
		if err != nil {
			return nil, err
		}

		if region == row[1] {

		}
		agencies = append(agencies, Agency{
			ID:                   int64(id),
			Name:                 row[2],
			Address:              row[3],
			PhoneNumber:          row[4],
			RegistrationDate:     time,
			AgencyEmployeesCount: int32(employeesCount),
		})
	}

	return agencies, nil
}

// func get(csvFile *os.File, id int64) (*Agency, error) {
// 	stringID := strconv.Itoa(int(id))

// 	rows, err := readCSVFile(csvFile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, row := range rows {
// 		if row[0] == stringID {

// 		}
// 	}
// }

// func parseAgencyFromCSV()
