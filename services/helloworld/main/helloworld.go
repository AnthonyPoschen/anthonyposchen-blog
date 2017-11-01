package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/zanven42/anthonyposchen-blog/services/helloworld"
	"google.golang.org/grpc"
)

// server implements the grpc service of helloworld.
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: in.Name + " and go + grpc"}, nil
}

// Handlers holds routes and variables utilised within the routes.
type handlers struct {
	// hot indicates that we will not serve static files but instead pipe the request
	// to npm to handle to enable hot reloading.
	hot string

	files packr.Box
}

func main() {
	hot := flag.String("hot", "", "all static content will be forwarded to the address provided if present i.e '-hot=:3000' | '-hot=localhost:3000'")

	flag.Parse()

	// create grpc server
	grpcServer := grpc.NewServer()
	// register our service onto the grpc server. (can register many services if needed)
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	// wrap our grpc server in grpc-web so the frontend can talk to it.
	wrappedServer := grpcweb.WrapServer(grpcServer)

	// setup our http server to have a custom handler so each request can first be checked if its
	// a grpc request before its passed off to the default http serve mux.
	httpServer := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if wrappedServer.IsGrpcWebRequest(r) {
				wrappedServer.ServeHTTP(w, r)
				return
			}
			http.DefaultServeMux.ServeHTTP(w, r)
		}),
		Addr: "localhost:8080",
	}
	if *hot != "" {
		fmt.Printf("Hot loading set to address: %s\n", *hot)
	}
	h := handlers{hot: *hot, files: packr.NewBox("./../../../frontend/dist")}
	http.HandleFunc("/", h.staticContent)
	fmt.Println("Listening on", httpServer.Addr)
	fmt.Println(httpServer.ListenAndServe())
}

func (h handlers) staticContent(w http.ResponseWriter, r *http.Request) {
	// SPA if statement to make sure the page is fetched which requires the app to self route.
	// remove this if statement if you are not running a SPA.
	if filepath.Ext(r.URL.Path) == "" {
		r.URL.Path = "/"
	}

	if h.hot == "" {
		http.FileServer(h.files).ServeHTTP(w, r)
		return
	}

	// become a proxy to forward the request / results.
	req, err := http.NewRequest(r.Method, "http://"+h.hot+r.URL.Path+r.URL.RawQuery, r.Body)
	if err != nil {
		fmt.Printf("Failed to build request %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for k, v := range r.Header {
		req.Header.Set(k, v[0])
	}

	client := http.Client{Timeout: time.Second * 2}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching from hot server %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("failed to extract result from hot result: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	for k, v := range res.Header {
		w.Header().Set(k, v[0])
	}
	w.Write(body)
}
