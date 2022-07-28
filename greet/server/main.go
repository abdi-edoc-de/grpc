package main

import (
	"log"
	"net"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

)

var address = "localhost:50051"
type server struct{
	pb.GreetServiceServer
}
func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil{
		log.Fatalf("failed to listen : %v", err)
	}
	log.Printf("listening on %v", address)
	tls := true
	opts := []grpc.ServerOption{}
	if tls{
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("error while loading certeficats : %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to serve")
	}
}