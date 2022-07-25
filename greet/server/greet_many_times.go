package main

import (
	"fmt"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)
func (s *server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked : %v", in)
	for i:=0; i < 10; i++{
		res := fmt.Sprintf("name : %v and number : %v", in.Name, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}