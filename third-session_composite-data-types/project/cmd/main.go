package main

import (
	"fmt"

	companyagencies "github.com/morteza-shahrabi-farahani/gocast-bootcamp"
)

func main() {
	csvFile, err := companyagencies.ConnectToCSV()
	if err != nil {
		fmt.Println("error in connecting to the csv file")
		return
	}
	defer csvFile.Close()

	firstTime := true
	err = companyagencies.GetCommands(csvFile, firstTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	firstTime = false
	for {
		err = companyagencies.GetCommands(csvFile, firstTime)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
