package handlers

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

func StartShiftWithClient(client pb.EmployeeControllerClient) http.HandlerFunc {
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
			_, err = w.Write([]byte(`{"status":"success"}`))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`{"code":"%d", "status":"%s"}`, code.GetCode(), code.GetError())))
		}
		if err != nil {
			Log.Printf("Error during writing message to writer [start shift], err:%s \n", err.Error())
			return
		}
	}

}

func EndShiftWithClient(client pb.EmployeeControllerClient) http.HandlerFunc {
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
			_, err = w.Write([]byte(`{"status":"success"}`))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`{"code":"%d", "status":"%s"}`, code.GetCode(), code.GetError())))
		}
		if err != nil {
			Log.Printf("Error during writing message to writer [end shift], err:%s \n", err.Error())
			return
		}
	}

}
