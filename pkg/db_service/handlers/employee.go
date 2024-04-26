package handlers

import (
	"context"
	"log"

	"github.com/ring0-rootkit/golang-staff-mngmnt/grpc"
	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/common/logging"
	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/db_service/repository"
)

var Log *log.Logger = logging.GetFor("db_service repository")

var (
	successCode          = grpc.ResponseCode{Code: 200}
	noIdAndNameErrorCode = grpc.ResponseCode{Code: 400, Error: "no 'id' or 'name + surname' was provided"}
)

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

func (EmployeeServer) GetWorkedHours(ctx context.Context, e *grpc.Employee) (*grpc.HoursWorked, error) {
	var id int64
	if e.GetId() != 0 {
		id = e.GetId()
	} else if e.GetName() != "" && e.GetSurname() != "" {
		name := e.GetName()
		surname := e.GetSurname()

		id = repository.EmployeeIdByName(name, surname)
	} else {
		return &grpc.HoursWorked{ResponseCode: &noIdAndNameErrorCode}, nil
	}

	h := repository.GetHoursWorked(id)
	if h == -1 {
		code := &grpc.ResponseCode{Code: 400, Error: "user id not found in database"}
		return &grpc.HoursWorked{Hours: h, ResponseCode: code}, nil
	}
	return &grpc.HoursWorked{Hours: h, ResponseCode: &successCode}, nil
}

func (EmployeeServer) GetSalaryPerHour(ctx context.Context, e *grpc.Employee) (*grpc.SalaryPH, error) {
	var id int64
	if e.GetId() != 0 {
		id = e.GetId()
	} else if e.GetName() != "" && e.GetSurname() != "" {
		name := e.GetName()
		surname := e.GetSurname()
		id = repository.EmployeeIdByName(name, surname)
	} else {
		return &grpc.SalaryPH{ResponseCode: &noIdAndNameErrorCode}, nil
	}

	s := repository.GetSalaryPerHour(id)
	return &grpc.SalaryPH{Salary: &grpc.Salary{Dollars: s.SpHDollars, Cents: s.SpHCents},
		ResponseCode: &successCode}, nil
}
