package rpc_handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "github.com/ring0-rootkit/golang-staff-mngmnt/grpc"
	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/common/logging"
)

var Log *log.Logger = logging.GetFor("middlware http handler")

func StartShift(client pb.EmployeeControllerClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			Log.Printf("Error while getting path value (id) from path: %s \n", r.URL.Path)
			return
		}

		code, err := client.StartWorkShift(context.Background(), &pb.Employee{Id: id})
		if err != nil {
			Log.Printf("Error when trying to send grpc call. Err:%s\n", err.Error())
			return
		}

		if code.GetCode() == 200 {
			_, err = w.Write([]byte(`"status":"success"`))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`{"code":"%d", "status":"%s"}`, code.GetCode(), code.GetError())))
		}
		if err != nil {
			Log.Printf("Error during handling response code [startShift], err:%s \n", err.Error())
		}
	}

}

func EndShift(client pb.EmployeeControllerClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			Log.Printf("Error while getting path value (id) from path: %s \n", r.URL.Path)
			return
		}

		code, err := client.EndWorkShift(context.Background(), &pb.Employee{Id: id})
		if err != nil {
			Log.Printf("Error when trying to send grpc call. Err:%s\n", err.Error())
			return
		}
		if code.GetCode() == 200 {
			_, err = w.Write([]byte(`"status":"success"`))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`{"code":"%d", "status":"%s"}`, code.GetCode(), code.GetError())))
		}
		if err != nil {
			Log.Printf("Error during handling response code [endShift], err:%s \n", err.Error())
		}
	}
}

func GetWorkedHours(client pb.EmployeeControllerClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			Log.Printf("Error while getting path value (id) from path: %s \n", r.URL.Path)
			return
		}

		hours, err := client.GetWorkedHours(context.Background(), &pb.Employee{Id: id})
		if err != nil {
			Log.Printf("Error when trying to send grpc call. Err:%s\n", err.Error())
			return
		}

		if hours.GetResponseCode().GetCode() == 200 {
			_, err = w.Write([]byte(fmt.Sprintf(`"status":"success", "hours":"%f"`, hours.GetHours())))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`{"code":"%d", "status":"%s"}`,
				hours.GetResponseCode().GetCode(),
				hours.GetResponseCode().GetError())))
		}
		if err != nil {
			Log.Printf("Error during handling response code [getWorkedHours], err:%s \n", err.Error())
		}
	}
}

func GetWorkedHoursByName(client pb.EmployeeControllerClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		surname := r.URL.Query().Get("surname")

		if name == "" || surname == "" {
			Log.Printf("no name or surname was provided at %s", r.URL.String())
			_, err := w.Write([]byte("no name or surname was provided"))
			if err != nil {
				Log.Printf("cannot write to RepsonseWriter at %s", r.URL.String())
			}
			return
		}

		hours, err := client.GetWorkedHours(context.Background(), &pb.Employee{Name: name, Surname: surname})
		if err != nil {
			Log.Printf("Error when trying to send grpc call. Err:%s\n", err.Error())
			return
		}

		if hours.GetResponseCode().GetCode() == 200 {
			_, err = w.Write([]byte(fmt.Sprintf(`"status":"success", "hours":"%f"`,
				hours.GetHours())))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`{"code":"%d", "status":"%s"}`,
				hours.GetResponseCode().GetCode(),
				hours.GetResponseCode().GetError())))
		}
		if err != nil {
			Log.Printf("Error during handling response code [getWorkedHoursByName], err:%s \n", err.Error())
		}
	}
}

func SalaryPerHour(client pb.EmployeeControllerClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			Log.Printf("Error while getting path value (id) from path: %s \n", r.URL.Path)
			return
		}

		salaryPH, err := client.GetSalaryPerHour(context.Background(), &pb.Employee{Id: id})
		if err != nil {
			Log.Printf("Error when trying to send grpc call. Err:%s\n", err.Error())
			return
		}

		if salaryPH.GetResponseCode().GetCode() == 200 {
			_, err = w.Write([]byte(fmt.Sprintf(`"status":"success", "salary":{"dollars":"%d", "cents":"%d"}`,
				salaryPH.GetSalary().GetDollars(),
				salaryPH.GetSalary().GetCents())))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`{"code":"%d", "status":"%s"}`,
				salaryPH.GetResponseCode().GetCode(),
				salaryPH.GetResponseCode().GetError())))
		}
		if err != nil {

			Log.Printf("Error during handling response code [getSalaryPerHour], err:%s \n", err.Error())
		}
	}

}

func SalaryPerHourByName(client pb.EmployeeControllerClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		surname := r.URL.Query().Get("surname")

		if name == "" || surname == "" {
			Log.Printf("no name or surname was provided at %s", r.URL.String())
			_, err := w.Write([]byte("no name or surname was provided"))
			if err != nil {
				Log.Printf("cannot write to RepsonseWriter at %s", r.URL.String())
			}
			return
		}

		salaryPH, err := client.GetSalaryPerHour(context.Background(), &pb.Employee{Name: name, Surname: surname})
		if err != nil {
			Log.Printf("Error when trying to send grpc call. Err:%s\n", err.Error())
			return
		}

		if salaryPH.GetResponseCode().GetCode() == 200 {
			_, err = w.Write([]byte(fmt.Sprintf(`"status":"success", "salary":{"dollars":"%d", "cents":"%d"}`,
				salaryPH.GetSalary().GetDollars(),
				salaryPH.GetSalary().GetCents())))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`{"code":"%d", "status":"%s"}`,
				salaryPH.GetResponseCode().GetCode(),
				salaryPH.GetResponseCode().GetError())))
		}
		if err != nil {

			Log.Printf("Error during handling response code [getSalaryPerHourByName], err:%s \n", err.Error())
		}
	}
}
