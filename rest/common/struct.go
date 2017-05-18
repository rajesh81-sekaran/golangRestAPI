package common

import (
	"fmt"
)

// ToSlice converts map of address book to slice returns []AddressBook
func ToSlice(abMap map[string]AddressBook) []AddressBook {
	if len(abMap) == 0 {
		return []AddressBook{}
	}
	abSlice := []AddressBook{}
	for _, v := range abMap {
		abSlice = append(abSlice, v)
	}
	return abSlice
}

// ToMap converts slice of address book to map returns map[string]AddressBook
func ToMap(abSlice []AddressBook) map[string]AddressBook {
	if len(abSlice) == 0 {
		err := MyError{errorString: "slice of address book is empty"}
		fmt.Printf("%s\n", err.Error())
		return map[string]AddressBook{}
	}
	abMap := map[string]AddressBook{}
	for _, v := range abSlice {
		abMap[v.FirstName] = v
	}
	return abMap
}
