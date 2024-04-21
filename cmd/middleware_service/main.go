package main

import (
	"net/http"

	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/middlware_service/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee/{id}/start", handlers.StartShift)
	mux.HandleFunc("/employee/{id}/end", handlers.EndShift)
	_ = http.ListenAndServe(":8080", mux)
}
