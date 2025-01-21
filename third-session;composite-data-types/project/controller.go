package companyagencies

import (
	"flag"
	"fmt"
	"os"
)

func getCommands(csvFile *os.File) error {
	command := flag.String("command", "list", "Please enter the command")
	region := flag.String("region", "Tehran", "This command specifies the city of the agency.")
	flag.Parse()

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
		// var id int64
		// fmt.Println("Please enter id of the agency")
		// fmt.Scanf("%d", &id)
		// agency, err := get(csvFile, id)
		// if err != nil {
		// 	return nil, err
		// }

		// fmt.Printf("%+v\n", agency)
	case "create":
		// create()
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
