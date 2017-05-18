# golangRestAPI
golang Rest API

Methods/Services:
With the current implementation the server will listen on port “9090”.
The following are the services exposed as part of the RestAPI implementation:
add: 
POST hostname:portnumber /add
E.g:
http://192.168.33.10:9090/add
Message body:
{"firstName":"FirstName",
"lastName":"LastName",
"email":"Email",
"phoneNumber":1111111111}

Service description:
Using add we can add an Address entry. The duplicates are not allowed. The “firstName” field is considered as the primary key and duplicates are not allowed.

show:
GET hostname:portnumber /show?name=FirstName0
E.g:
http://192.168.33.10:9090/show?name=FirstName0

Service description:
Using  show we can display an Address entry.  The “name” key should be used to fetch the Address entry, the “firstName” that macthes the value of “name” key will be retrieved.

list:
GET hostname:portnumber /list
E.g:
http://192.168.33.10:9090/list

Service description:
Using  list all the Address entries will be retrieved. 

modify:
POST hostname:portnumber /modify

E.g:
http://192.168.33.10:9090/modify
Message body:
{"firstName":"FirstName22",
"lastName":"LastName1",
"email":"",
"phoneNumber":5713761273
}

Service description:
Using  modify an Address entry that matches the provided “firstName” will be modified.

delete:
GET hostname:portnumber /delete?name=FirstName0

E.g
http://192.168.33.10:9090/delete?name=FirstName0

Service description:
Using  delete we can delete an Address entry.  The “name” key should be used to delete the Address entry, the “firstName” that macthes the value of “name” key will be deleted.

save:
GET/POST hostname:portnumber /save

E.g:
http://192.168.33.10:9090/save

Service description:
Using  save all Address entries in that are stored as map in the server buffer will be stored to a csv file. This csv file can be used to load the values when the server is restarted. It makes a persistent data store.

load:
GET/POST hostname:portnumber /load

E.g:
http://192.168.33.10:9090/load

Service description:
Using  load the Address entries in that are stored in the csv file can be loaded as map in the server buffer. This csv file can be used to load the values when the server is restarted. It makes a persistent data store.

External dependencies:
To make the server design simple I did not use much of external dependencies and used csv as the data store. We can use any data base instead of the csv file.
I have used the following csv package for golang structure conversion to csv file
“github.com/gocarina/gocsv”

Address entry structure in golang:
type AddressBook struct {
    FirstName   string `json:"firstName" csv:"firstName"`
    LastName    string `json:"lastName" csv:"lastName"`
    Email       string `json:"email" csv:"email"`
    PhoneNumber int    `json:"phoneNumber" csv:"phoneNumber"`
}



Address entry csv file in the filesystem:
firstName	lastName	email	phoneNumber
FirstName1	LastName1	Email1	111
FirstName2	LastName2	Email2	222
FirstName3	LastName3	Email3	333

Testing:
I have used postman and curl to exercise these APIs. I am working on the go test cases.

I have used enough locking mechanism using Locks to ensure data consistency considering parallel access from multiple clients.


