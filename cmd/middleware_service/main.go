package main

import (
	"net/http"

	"github.com/ring0-rootkit/golang-staff-mngmnt/internal/http_handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee/{id}/start", http_handler.StartShiftHandler)

	_ = http.ListenAndServe(":8080", mux)
}
