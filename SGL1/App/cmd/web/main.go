package main

import (
	"fmt"
	"myapp/pkg"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", pkg.HomePage).Methods("GET")
	r.HandleFunc("/maincharacters/{name}", pkg.MainChar).Methods("GET")
	r.HandleFunc("/healthcheck/", pkg.HealthCheck).Methods("GET")

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}
