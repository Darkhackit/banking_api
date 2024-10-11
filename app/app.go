package app

import (
	"fmt"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", Greet)
	mux.HandleFunc("/customers", GetAllCustomers)

	// Start the server and listen on localhost:8080
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
