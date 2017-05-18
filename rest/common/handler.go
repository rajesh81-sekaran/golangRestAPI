package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleAdd Handler function for add service
func HandleAdd(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("In handleAdd")
	var ab AddressBook
	if r.Method != http.MethodPost {
		err := fmt.Sprintf("Method %s not supported on this action %s\n", r.Method, r.URL.Path)
		fmt.Printf("%s", err)
		http.Error(w, err, http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&ab); err != nil {

		msg := fmt.Sprintf("Error %s while decoding json\n", err.Error())
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	fmt.Printf("Adding %s to the address book\n", ab.FirstName)
	Mutex.Lock()
	defer Mutex.Unlock()
	if _, ok := AddrBook[ab.FirstName]; ok {
		AddrBook[ab.FirstName] = ab
		msg := fmt.Sprintf("Overwritten name %s present in the address book\n", ab.FirstName)
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusContinue)

	} else {
		AddrBook[ab.FirstName] = ab
		msg := fmt.Sprintf("Added name %s present in the address book\n", ab.FirstName)
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusOK)
	}

}

// HandleModify Handler function for modify service
func HandleModify(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("In handleModify")
	var ab AddressBook
	if r.Method != http.MethodPost {
		err := fmt.Sprintf("Method %s not supported on this action %s\n", r.Method, r.URL.Path)
		fmt.Printf("%s", err)
		http.Error(w, err, http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&ab); err != nil {
		msg := fmt.Sprintf("Error %s while decoding json\n", err.Error())
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if ab.FirstName == "" {
		msg := "Name not provided as part of the query\n"
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	Mutex.RLock()
	defer Mutex.RUnlock()
	if len(AddrBook) == 0 {
		msg := fmt.Sprintf("Address book is empty nothing to modify\n")
		fmt.Printf("%s\n", msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	if abTmp, ok := AddrBook[ab.FirstName]; !ok {
		msg := fmt.Sprintf("%s not found in the Address book, nothing to modify\n", abTmp.FirstName)
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	abTmp, _ := AddrBook[ab.FirstName]
	if ab.LastName != "" && abTmp.LastName != ab.LastName {
		abTmp.LastName = ab.LastName
	}
	if ab.Email != "" && abTmp.Email != ab.Email {
		abTmp.Email = ab.Email
	}
	if ab.PhoneNumber != 0 && abTmp.PhoneNumber != ab.PhoneNumber {
		abTmp.PhoneNumber = ab.PhoneNumber
	}
	AddrBook[ab.FirstName] = abTmp
	msg := fmt.Sprintf("Modified name %s present in the address book\n", ab.FirstName)
	fmt.Printf("%s", msg)
	http.Error(w, msg, http.StatusOK)
}

// HandleDelete Handler function for delete service
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("In handleDelete")
	if r.Method != http.MethodGet {
		msg := fmt.Sprintf("Method %s not supported on this action %s\n", r.Method, r.URL.Path)
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}
	Mutex.RLock()
	defer Mutex.RUnlock()
	if len(AddrBook) == 0 {
		msg := fmt.Sprintf("Address book is empty\n")
		fmt.Printf("%s\n", msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	firstName := r.URL.Query().Get("name")
	if firstName == "" {
		msg := "Name not provided as part of the query\n"
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	if _, ok := AddrBook[firstName]; ok {
		delete(AddrBook, firstName)
		msg := fmt.Sprintf("%s deleted from the address book", firstName)
		fmt.Printf("%s\n", msg)
		http.Error(w, msg, http.StatusOK)
	} else {
		msg := fmt.Sprintf("%s not found in the address book", firstName)
		fmt.Printf("%s\n", msg)
		http.Error(w, msg, http.StatusNotFound)
	}
}

// HandleShow Handler function for show service
func HandleShow(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("In handleShow")
	if r.Method != http.MethodGet {
		msg := fmt.Sprintf("Method %s not supported on this action %s\n", r.Method, r.URL.Path)
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}
	Mutex.RLock()
	defer Mutex.RUnlock()
	if len(AddrBook) == 0 {
		msg := fmt.Sprintf("Address book is empty\n")
		fmt.Printf("%s\n", msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	firstName := r.URL.Query().Get("name")
	if firstName == "" {
		msg := "Name not provided as part of the query\n"
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	if ab, ok := AddrBook[firstName]; ok {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ab); err != nil {
			msg := fmt.Sprintf("Error %s while encoding json\n", err.Error())
			fmt.Printf("%s\n", msg)
			http.Error(w, msg, http.StatusInternalServerError)
		}
	} else {
		msg := fmt.Sprintf("%s not found in the address book", firstName)
		fmt.Printf("%s\n", msg)
		http.Error(w, msg, http.StatusNotFound)
	}
}

// HandleList Handler function for list service
func HandleList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("In handleList")
	Mutex.RLock()
	defer Mutex.RUnlock()
	if len(AddrBook) == 0 {
		msg := fmt.Sprintf("Address book is empty\n")
		fmt.Printf("%s\n", msg)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	for _, ab := range AddrBook {
		if err := json.NewEncoder(w).Encode(ab); err != nil {
			msg := fmt.Sprintf("Error %s while encoding json\n", err.Error())
			fmt.Printf("%s\n", msg)
			http.Error(w, msg, http.StatusInternalServerError)
		}
	}
}

// HandleSave Handler function for save service
func HandleSave(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	Mutex.Lock()
	defer Mutex.Unlock()
	fmt.Println("In handleSave")

	if len(AddrBook) == 0 {
		msg := fmt.Sprintf("Saving empty address book\n")
		fmt.Printf("%s\n", msg)
	}
	if err := WriteToCSV(CSVFile, AddrBook); err != nil {
		msg := fmt.Sprintf("Error %s after writeToCSV\n", err.Error())
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	msg := fmt.Sprintf("Saved address book to file")
	fmt.Printf("%s\n", msg)
	http.Error(w, msg, http.StatusOK)
}

// HandleLoad Handler function for load service
func HandleLoad(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("In handleLoad")
	Mutex.RLock()
	defer Mutex.RUnlock()
	abMap, err := ReadFromCSV(CSVFile)
	if err != nil {
		msg := fmt.Sprintf("%s after ReadFromCSV\n", err.Error())
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	AddrBook = abMap
	msg := fmt.Sprintf("Loaded address book file")
	fmt.Printf("%s\n", msg)
	http.Error(w, msg, http.StatusOK)
}
