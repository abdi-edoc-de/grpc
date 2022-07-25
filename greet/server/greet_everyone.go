package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)


func (s *server) GreetEveryone( stream pb.GreetService_GreetEveryoneServer) error{
	log.Printf("GreetEveryone function was invoked ")
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			log.Fatalf("error while reading stream from client in server : %v", err)
		}
		res := fmt.Sprintf("Hello %v !! \n", req.Name)
		
		if err := stream.Send(&pb.GreetResponse{Result: res,}); err!= nil{
			log.Fatalf("error while sengin response : %v", err)
		}
	}
	return nil
}