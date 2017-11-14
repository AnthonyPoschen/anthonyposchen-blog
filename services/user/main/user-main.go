package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/zanven42/anthonyposchen-blog/services/user"
	"github.com/zanven42/anthonyposchen-blog/util"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	env := util.GetENV()
	port := env.Services.User.Port

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})
	fmt.Println(grpcServer.Serve(lis))
}
