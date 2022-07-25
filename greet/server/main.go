package main

import (
	"log"
	"net"
	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
	"google.golang.org/grpc"
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
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to serve")
	}
}