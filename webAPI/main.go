package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Creating a slice of employee objects
var company []employee

//Creating an employee struct to hold employee data
type employee struct {
	ID      int    `json:"Id"`
	Name    string `json:"Name"`
	Role    string `json:"Role"`
	Address string `json:"Address"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnData")
	json.NewEncoder(w).Encode(company)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/employeeData", returnData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	//Employee objects (dummy data for testing)

	company = []employee{
		employee{
			ID:      1,
			Name:    "Sherlock Holmes",
			Role:    "UI Designer",
			Address: "221B Baker Street"},
		employee{
			ID:      2,
			Name:    "Sirius Black",
			Role:    "Project Manager",
			Address: "12 Grimmauld Place"}}

	handleRequests()
}
