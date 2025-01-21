package companyagencies

import "time"

type Agency struct {
	ID                   int64
	Region               string
	Name                 string
	Address              string
	PhoneNumber          string
	RegistrationDate     time.Time
	AgencyEmployeesCount int32
}
