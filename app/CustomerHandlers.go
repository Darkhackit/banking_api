package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"github/Darkhackit/banking_api/service"
	"net/http"
)

//	type Customer struct {
//		Name    string `json:"full_name" xml:"full_name"`
//		City    string `json:"city" xml:"city"`
//		ZipCode string `json:"zip_code" xml:"zip_code"`
//	}
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		return
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			_, err := fmt.Fprintf(w, err.Error())
			if err != nil {
				return
			}
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			_, err := fmt.Fprintf(w, err.Error())
			if err != nil {
				return
			}
		}
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer, err := ch.service.GetCustomerById(vars["customer_id"])
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)

	}
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}
