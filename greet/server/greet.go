package main

import (
	"context"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error){
	log.Printf("Greet function invoked with %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.Name,
	}, nil
}