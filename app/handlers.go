package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello World!")
	if err != nil {
		return
	}
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Emma", "Takoradi", "2839"},
		{"Derby", "Accra", "2039"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}

	}
}
