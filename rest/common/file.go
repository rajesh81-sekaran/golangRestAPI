package common

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

// WriteToCSV Read address book data from map of address book and write to csv file
func WriteToCSV(abFileName string, abMap map[string]AddressBook) error {
	if len(abMap) == 0 {
		msg := fmt.Sprintf("address book is empty while writing to csv file")
		fmt.Printf("%s\n", msg)
	}
	abFile, err := os.OpenFile(abFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		msg := fmt.Sprintf("Error %s while opening the address book file", err.Error())
		fmt.Printf("%s", msg)
		return MyError{errorString: msg}
	}
	abSlice := ToSlice(abMap)
	err = gocsv.MarshalFile(&abSlice, abFile)
	abFile.Sync()
	abFile.Close()
	if err != nil {
		msg := fmt.Sprintf("Error %s while marshalling address book to csv file\n", err.Error())
		fmt.Printf("%s", msg)
		return MyError{errorString: msg}
	}
	msg := fmt.Sprintf("Successfully saved the address book to csv file\n")
	fmt.Printf("%s", msg)
	return nil
}

// ReadFromCSV Read address book data from csv file and return the map of address book
func ReadFromCSV(abFileName string) (map[string]AddressBook, error) {
	if abFileName == "" {
		msg := fmt.Sprintf("address book file name is empty while reading from csv")
		fmt.Printf("%s", msg)
		return nil, MyError{errorString: msg}
	}
	abFile, err := os.OpenFile(abFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		msg := fmt.Sprintf("Error %s while opening the address book file for reading", err.Error())
		fmt.Printf("%s", msg)
		return nil, MyError{errorString: msg}
	}
	defer abFile.Close()
	abSlice := []AddressBook{}

	if stat, err := abFile.Stat(); err != nil {
		msg := fmt.Sprintf("Error %s", err.Error())
		fmt.Printf("%s\n", msg)
		return nil, MyError{errorString: msg}
	} else if stat.Size() == 0 {
		fmt.Printf("Input csv File is empty\n")
		return map[string]AddressBook{}, nil
	}

	if err := gocsv.UnmarshalFile(abFile, &abSlice); err != nil {
		msg := fmt.Sprintf("Error %s while unmarshalling address book csv file to address book slice", err.Error())
		fmt.Printf("%s\n", msg)
		return nil, MyError{errorString: msg}
	}

	abMap := ToMap(abSlice)
	msg := fmt.Sprintf("Successfully read data from address book to csv\n")
	fmt.Printf("%s", msg)
	return abMap, nil
}
