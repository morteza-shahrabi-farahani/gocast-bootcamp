package companyagencies

import (
	"flag"
	"fmt"
	"os"
)

func GetCommands(csvFile *os.File, isFistTime bool) error {
	var command, region *string
	var newCommand, newRegion string

	if isFistTime {
		command = flag.String("command", "list", "Please enter the command")
		region = flag.String("region", "Tehran", "This command specifies the city of the agency.")
		flag.Parse()
	} else {
		fmt.Println("Please enter the command")
		fmt.Scanf("%s", &newCommand)
		command = &newCommand
		fmt.Println("Please enter the region")
		fmt.Scanf("%s", &newRegion)
		region = &newRegion
	}

	switch *command {

	case "list":
		agencies, err := list(csvFile, *region)
		if err != nil {
			return err
		}

		for _, agency := range agencies {
			fmt.Printf("%+v\n", agency)
		}

	case "get":
		var id int64
		fmt.Println("Please enter id of the agency")
		fmt.Scanf("%d", &id)

		agency, err := get(csvFile, id)
		if err != nil {
			return err
		}

		fmt.Printf("%+v\n", agency)

	case "create":
		var name string
		fmt.Println("Please enter name of the agency")
		fmt.Scanf("%s", &name)

		var address string
		fmt.Println("Please enter address of the agency")
		fmt.Scanf("%s", &address)

		var phoneNumber string
		fmt.Println("Please enter phone number of the agency")
		fmt.Scanf("%s", &phoneNumber)

		var AgencyEmployeesCount int32
		fmt.Println("Please enter number of agency employees")
		fmt.Scanf("%d", &AgencyEmployeesCount)

		agency := &Agency{
			Region:               *region,
			Name:                 name,
			Address:              address,
			PhoneNumber:          phoneNumber,
			AgencyEmployeesCount: AgencyEmployeesCount,
		}

		agencyID, err := create(csvFile, agency)
		if err != nil {
			return err
		}

		fmt.Println("id of the inserted agency is ", agencyID)

	case "edit":
		// edit()
	case "status":
		// status()
	case "exit":
		// exit()
	default:
		fmt.Println("Please enter the valid command")
	}

	return nil
}
