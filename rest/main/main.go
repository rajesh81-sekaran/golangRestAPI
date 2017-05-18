package main

import (
	"fmt"
	"net/http"
	cmn "rest/common"
	"sync"
)

func main() {
	cmn.AddrBook = map[string]cmn.AddressBook{}
	cmn.Mutex = sync.RWMutex{}

	http.HandleFunc("/add", cmn.HandleAdd)
	http.HandleFunc("/modify", cmn.HandleModify)
	http.HandleFunc("/delete", cmn.HandleDelete)
	http.HandleFunc("/show", cmn.HandleShow)
	http.HandleFunc("/list", cmn.HandleList)
	http.HandleFunc("/save", cmn.HandleSave)
	http.HandleFunc("/load", cmn.HandleLoad)
	fmt.Println("In main() starting server")
	http.ListenAndServe(":"+cmn.Port, nil)
}
