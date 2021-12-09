package main

import (
	"fmt"
    "net/http"
    "log"

    "github.com/gorilla/mux"
)


type Tag struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func getTags(w http.ResponseWriter, r *http.Request) {
	results, err := db.Query("SELECT tag_id, name FROM tag")

    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        fmt.Fprintf(w, tag.Name)
    }
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/tags", getTags)
    
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}