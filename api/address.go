package handler

import (
	"encoding/json"
	"net/http"
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func AddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      1993,
		Phone:     "0870940955",
	}
	json.NewEncoder(w).Encode(addBook)
}
