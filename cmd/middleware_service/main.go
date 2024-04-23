package main

import (
	"fmt"
	"net/http"

	pb "github.com/ring0-rootkit/golang-staff-mngmnt/grpc"
	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/middlware_service/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

	mux.HandleFunc("/employee/{id}/start", handlers.StartShiftWithClient(client))
	mux.HandleFunc("/employee/{id}/end", handlers.EndShiftWithClient(client))

	_ = http.ListenAndServe(":8080", mux)
}
