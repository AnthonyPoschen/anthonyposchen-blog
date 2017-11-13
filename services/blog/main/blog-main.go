package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/zanven42/anthonyposchen-blog/services/blog"

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
	// start grpc web host here

	// setup the blog service for remote access.
	// only from internal access points.

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, &server{})
	fmt.Println(http.ListenAndServe(":8082", grpcServer))
}
