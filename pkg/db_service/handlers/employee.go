package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ring0-rootkit/golang-staff-mngmnt/grpc"
	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/common/logging"
	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/db_service/repository"
)

var Log *log.Logger = logging.GetFor("db_service repository")

type EmployeeServer struct {
	grpc.UnimplementedEmployeeControllerServer
}

func (EmployeeServer) StartWorkShift(ctx context.Context, e *grpc.Employee) (*grpc.ResponseCode, error) {
	Log.Printf("Got employee for start workshift id:%d", e.GetId())
	return &grpc.ResponseCode{Code: 200, Error: ""}, nil
}

func (EmployeeServer) EndWorkShift(ctx context.Context, e *grpc.Employee) (*grpc.ResponseCode, error) {
	Log.Printf("Got employee for end workshift id:%d", e.GetId())
	return &grpc.ResponseCode{Code: 200, Error: ""}, nil
}

func (EmployeeServer) GetWorkedHours(w http.ResponseWriter, r *http.Request) {
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

func (EmployeeServer) GetWorkedHoursByName(w http.ResponseWriter, r *http.Request) {
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

func (EmployeeServer) SalaryPerHour(w http.ResponseWriter, r *http.Request) {
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

func (EmployeeServer) SalaryPerHourByName(w http.ResponseWriter, r *http.Request) {
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
