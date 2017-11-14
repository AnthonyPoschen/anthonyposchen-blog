package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/zanven42/anthonyposchen-blog/services/blog"
	"github.com/zanven42/anthonyposchen-blog/util"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetBlog(ctx context.Context, in *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
	return &pb.GetBlogResponse{Blog: nil, Error: "Not implemented"}, fmt.Errorf("Not Implemented")
}

func (s *server) GetMultiBlog(ctx context.Context, in *pb.GetMultiBlogRequest) (*pb.GetMultiBlogResponse, error) {
	return &pb.GetMultiBlogResponse{Blogs: nil, Error: "Not implemented"}, fmt.Errorf("Not Implemented")
}

func (s *server) CreateBlog(ctx context.Context, in *pb.CreateBlogRequest) (*pb.CreateBlogResult, error) {
	return &pb.CreateBlogResult{Id: "", Error: "Not implemented"}, fmt.Errorf("Not Implemented")
}

func (s *server) DeleteBlog(ctx context.Context, in *pb.DelBlogRequest) (*pb.DelBlogResult, error) {
	return &pb.DelBlogResult{Ok: false, Error: "Not implemented"}, fmt.Errorf("Not Implemented")
}

func main() {
	env := util.GetENV()
	port := env.Services.Blog.Port

	// start grpc web host here
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, &server{})
	fmt.Println(grpcServer.Serve(lis))
}
