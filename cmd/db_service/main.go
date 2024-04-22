package main

import (
	"net/http"

	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/db_service/handlers"
)

//TODO: add admin page

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/employee/{id}/hours", handlers.GetWorkedHours)
	mux.HandleFunc("/employee/hours", handlers.GetWorkedHoursByName)
	mux.HandleFunc("/employee/{id}/salary", handlers.SalaryPerHour)
	mux.HandleFunc("/employee/salary", handlers.SalaryPerHourByName)

	_ = http.ListenAndServe(":9000", mux)
}
