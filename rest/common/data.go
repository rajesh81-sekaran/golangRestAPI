package common

import "sync"

// MyError structure
type MyError struct {
	errorString string
}

func (me MyError) Error() string {
	return me.errorString
}

//AddressBook structure
type AddressBook struct {
	FirstName   string `json:"firstName" csv:"firstName"`
	LastName    string `json:"lastName" csv:"lastName"`
	Email       string `json:"email" csv:"email"`
	PhoneNumber int    `json:"phoneNumber" csv:"phoneNumber"`
}

//Mutex Lock used across the different handlers for synchronization
var Mutex sync.RWMutex

//AddrBook map of addressBook as map[string]AddressBook
var AddrBook map[string]AddressBook

// Port port on which the server listens
const Port string = "9090"

// CSVFile csv file name where the address will be stored
const CSVFile string = "./address.csv"
