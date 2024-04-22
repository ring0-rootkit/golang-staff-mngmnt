package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/common/logging"
	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/db_service/repository"
)

var Log *log.Logger = logging.GetFor("db_service repository")

func GetWorkedHours(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(500)
		_, err = w.Write([]byte(`"error":"cannot parse an employee id"`))
		if err != nil {
			Log.Printf("error while trying to send response, err:%s", err.Error())
		}
		return
	}
	h := repository.GetHoursWorked(id)
	jsonResponse := fmt.Sprintf(`"hours":"%f"`, h)
	_, err = w.Write([]byte(jsonResponse))
	if err != nil {
		Log.Printf("error while trying to send response, err:%s", err.Error())
	}
}

func GetWorkedHoursByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	if name == "" || surname == "" {
		w.WriteHeader(500)
		_, err := w.Write([]byte(`"error":"cannot parse an employee name/surname"`))
		if err != nil {
			Log.Printf("error while trying to send response, err:%s", err.Error())
		}
		return
	}

	id := repository.EmployeeIdByName(name, surname)

	h := repository.GetHoursWorked(id)
	jsonResponse := fmt.Sprintf(`"hours":"%f"`, h)
	_, err := w.Write([]byte(jsonResponse))
	if err != nil {
		Log.Printf("error while trying to send response, err:%s", err.Error())
	}
}

func SalaryPerHour(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(500)
		_, err = w.Write([]byte(`"error":"cannot parse an employee id"`))
		if err != nil {
			Log.Printf("error while trying to send response, err:%s", err.Error())
		}
		return
	}
	h := repository.GetSalaryPerHour(id)
	jsonResponse := fmt.Sprintf(`"slaryPerHour":"%f"`, h)
	_, err = w.Write([]byte(jsonResponse))
	if err != nil {
		Log.Printf("error while trying to send response, err:%s", err.Error())
	}
}

func SalaryPerHourByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	if name == "" || surname == "" {
		w.WriteHeader(500)
		_, err := w.Write([]byte(`"error":"cannot parse an employee name/surname"`))
		if err != nil {
			Log.Printf("error while trying to send response, err:%s", err.Error())
		}
		return
	}
	id := repository.EmployeeIdByName(name, surname)

	h := repository.GetSalaryPerHour(id)
	jsonResponse := fmt.Sprintf(`"slaryPerHour":"%f"`, h)
	_, err := w.Write([]byte(jsonResponse))
	if err != nil {
		Log.Printf("error while trying to send response, err:%s", err.Error())
	}
}
