package main

import (
	"log"
	"net"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{
	pb.CalculatorServiceServer
}
var address = "0.0.0.0:50051"
func main(){
	lis, err := net.Listen("tcp", address)
	if err != nil{
		log.Fatalf("failed to listen : %v", err)
	}
	log.Printf("listening on %v", address)
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to serve : %v", err)
	}
}