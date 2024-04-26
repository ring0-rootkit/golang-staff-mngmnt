package main

import (
	"fmt"
	"net/http"

	pb "github.com/ring0-rootkit/golang-staff-mngmnt/grpc"
	rpc_handler "github.com/ring0-rootkit/golang-staff-mngmnt/pkg/middlware_service/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO implement GetWorkedHours and GetSalaryPerHour in gRPC client

func main() {
	// gRPC configuration and startup
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:42069", opts...)
	if err != nil {
		panic(fmt.Sprintf("can't start gRPC client, err:%s", err.Error()))
	}
	defer conn.Close()

	client := pb.NewEmployeeControllerClient(conn)
	// end of gRPC configuration and startup

	mux := http.NewServeMux()

	mux.HandleFunc("/employee/{id}/start", rpc_handler.StartShift(client))
	mux.HandleFunc("/employee/{id}/end", rpc_handler.EndShift(client))

	mux.HandleFunc("/employee/{id}/hours", rpc_handler.GetWorkedHours(client))
	mux.HandleFunc("/employee/hours", rpc_handler.GetWorkedHoursByName(client))
	mux.HandleFunc("/employee/{id}/salary", rpc_handler.SalaryPerHour(client))
	mux.HandleFunc("/employee/salary", rpc_handler.SalaryPerHourByName(client))

	_ = http.ListenAndServe(":8080", mux)
}
