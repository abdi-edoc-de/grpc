package main

import (
	"context"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
)


func (s *server) Sum(ctx context.Context,in  *pb.SumRequest) (*pb.SumResponse, error){
	log.Printf("sum was invoked with : %v", in)
	return &pb.SumResponse{
		Result: in.X + in.Y,
	}, nil
}

