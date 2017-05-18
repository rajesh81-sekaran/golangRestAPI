package common

import (
	"os"
	"reflect"
	"testing"
)

var sliceOfAddr1 = []AddressBook{
	AddressBook{FirstName: "firstName1", LastName: "lastName1", Email: "email1", PhoneNumber: 1111111111},
	AddressBook{FirstName: "firstName2", LastName: "lastName2", Email: "email2", PhoneNumber: 2222222222},
}

var mapOfAddr1 = map[string]AddressBook{
	"firstName1": AddressBook{FirstName: "firstName1", LastName: "lastName1", Email: "email1", PhoneNumber: 1111111111},
	"firstName2": AddressBook{FirstName: "firstName2", LastName: "lastName2", Email: "email2", PhoneNumber: 2222222222},
}

var sliceOfAddr2 = []AddressBook{
	AddressBook{FirstName: "firstName3", LastName: "lastName3", Email: "email3", PhoneNumber: 3333333333},
	AddressBook{FirstName: "firstName4", LastName: "lastName4", Email: "email4", PhoneNumber: 4444444444},
}

var mapOfAddr2 = map[string]AddressBook{
	"firstName3": AddressBook{FirstName: "firstName3", LastName: "lastName3", Email: "email3", PhoneNumber: 3333333333},
	"firstName4": AddressBook{FirstName: "firstName4", LastName: "lastName4", Email: "email4", PhoneNumber: 4444444444},
}

var sliceOfAddr3 = []AddressBook{
	AddressBook{FirstName: "firstName5", LastName: "lastName5", Email: "email5", PhoneNumber: 5555555555},
	AddressBook{FirstName: "firstName6", LastName: "lastName6", Email: "email6", PhoneNumber: 6666666666},
}

var mapOfAddr3 = map[string]AddressBook{
	"firstName5": AddressBook{FirstName: "firstName5", LastName: "lastName5", Email: "email5", PhoneNumber: 5555555555},
	"firstName6": AddressBook{FirstName: "firstName6", LastName: "lastName6", Email: "email6", PhoneNumber: 6666666666},
}

var toSliceTests = []struct {
	input    map[string]AddressBook
	expected []AddressBook
}{
	{mapOfAddr1, sliceOfAddr1},
	{mapOfAddr2, sliceOfAddr2},
	{mapOfAddr3, sliceOfAddr3},
}

var toMapTests = []struct {
	input    []AddressBook
	expected map[string]AddressBook
}{
	{sliceOfAddr1, mapOfAddr1},
	{sliceOfAddr2, mapOfAddr2},
	{sliceOfAddr3, mapOfAddr3},
}

var mapAddr = []map[string]AddressBook{
	mapOfAddr1,
	mapOfAddr2,
	mapOfAddr3,
}
var fileNames = []string{
	"tst1.csv",
	"tst2.csv",
	"tst3.csv",
}

func TestToSlice(t *testing.T) {
	ok := false
	for _, v := range toSliceTests {
		got := ToSlice(v.input)
		if reflect.TypeOf(got) != reflect.TypeOf(v.expected) || len(got) != len(v.expected) {
			t.Error("ToSlice failed for INPUT\n", v.input, "\n\nEXPECTED\n", v.expected, "\n\nGOT\n", got)
			return
		}
		for _, v1 := range got {
			ok = false
			for _, v2 := range v.expected {
				if reflect.DeepEqual(v1, v2) {
					ok = true
				}
			}
		}
		if !ok {
			t.Error("ToSlice failed for INPUT\n", v.input, "\n\nEXPECTED\n", v.expected, "\n\nGOT\n", got)
		}
	}
}

func TestToMap(t *testing.T) {
	ok := false
	for _, v := range toMapTests {
		got := ToMap(v.input)
		if reflect.TypeOf(got) != reflect.TypeOf(v.expected) || len(got) != len(v.expected) {
			t.Error("ToMap failed for INPUT\n", v.input, "\n\nEXPECTED\n", v.expected, "\n\nGOT\n", got)
			return
		}
		for _, v1 := range got {
			ok = false
			for _, v2 := range v.expected {
				if reflect.DeepEqual(v1, v2) {
					ok = true
				}
			}
		}
		if !ok {
			t.Error("ToMap failed for INPUT\n", v.input, "\n\nEXPECTED\n", v.expected, "\n\nGOT\n", got)
		}
	}
}

func TestWriteToCSV(t *testing.T) {
	for k, v := range mapAddr {
		err := WriteToCSV(fileNames[k], v)
		if err != nil {
			t.Error("TestWriteToCSV failed for INPUT\n", v, "\n\nEXPECTED\n", nil, "\n\nGOT\n", err.Error())
		}
	}

	for _, fileName := range fileNames {
		os.Remove(fileName)
	}
}

func TestReadFromCSV(t *testing.T) {
	for k, v := range mapAddr {
		err := WriteToCSV(fileNames[k], v)
		if err != nil {
			t.Error("TestReadFromCSV failed while trying to create input files for the test for INPUT\n", v, "\n\nEXPECTED\n", nil, "\n\nGOT\n", err.Error())
		}
	}

	for _, fileName := range fileNames {
		_, err := ReadFromCSV(fileName)
		if err != nil {
			t.Error("ReadFromCSV failed for INPUT\n", fileName, "\n\nEXPECTED\n", nil, "\n\nGOT\n", err.Error())
		}
	}

	for _, fileName := range fileNames {
		os.Remove(fileName)
	}
}
