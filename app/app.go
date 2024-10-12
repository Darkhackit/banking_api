package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github/Darkhackit/banking_api/domain"
	"github/Darkhackit/banking_api/service"
	"net/http"
)

func Start() {
	//mux := http.NewServeMux()
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers", CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods("GET")

	// Start the server and listen on localhost:8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
