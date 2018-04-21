package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define path and what function gets called for it
	http.HandleFunc("/", rootHandler)
	//set fatal logging
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// read requests and write response
func rootHandler(w http.ResponseWriter, r *http.Request) {
	//Reading URL path as the %s variable
	fmt.Fprintf(w, "working? %s\n", r.URL.Path[1:])
}
