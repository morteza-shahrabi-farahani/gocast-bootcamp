package companyagencies

import "time"

type Agency struct {
	ID                   int64
	Name                 string
	Address              string
	PhoneNumber          string
	RegistrationDate     time.Time
	AgencyEmployeesCount int32
}
