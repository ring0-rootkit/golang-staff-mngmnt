package main

import (
	"net/http"

	"github.com/ring0-rootkit/golang-staff-mngmnt/internal/db_service/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/employee/{id}/hours", handlers.WorkedHours)

	_ = http.ListenAndServe(":9000", mux)
}
