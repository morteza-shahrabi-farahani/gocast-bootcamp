package companyagencies

import "fmt"

func main() {
	csvFile, err := connectToCSV()
	if err != nil {
		fmt.Println("error in connecting to the csv file")
		return
	}

	for {
		err = getCommands(csvFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
