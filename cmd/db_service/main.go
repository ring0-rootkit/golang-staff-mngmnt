package main

import (
	"fmt"
	"net"

	pb "github.com/ring0-rootkit/golang-staff-mngmnt/grpc"
	"github.com/ring0-rootkit/golang-staff-mngmnt/pkg/db_service/handlers"
	"google.golang.org/grpc"
)

// TODO: add admin page
const port int32 = 42069

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		panic(fmt.Sprintf("can't start gRPC server, err:%s", err.Error()))
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterEmployeeControllerServer(grpcServer, handlers.EmployeeServer{})

	_ = grpcServer.Serve(lis)
}
