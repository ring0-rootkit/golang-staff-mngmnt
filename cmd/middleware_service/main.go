package main

import (
	"net/http"

	"github.com/ring0-rootkit/golang-staff-mngmnt/internal/middlware_service/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee/{id}/start", handlers.StartShiftHandler)
	mux.HandleFunc("/employee/{id}/end", handlers.EndShiftHandler)
	_ = http.ListenAndServe(":8080", mux)
}
