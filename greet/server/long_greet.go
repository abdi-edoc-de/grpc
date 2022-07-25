package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)

func (s *server ) LongGreet(stream pb.GreetService_LongGreetServer) error{
	log.Printf("LongGreet function was invoked with : %v", stream)
	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("error while reading from the stream in server : %v", err)
		}
		res += fmt.Sprintf("Hello, %v!!!\n", req.Name)

	}

}