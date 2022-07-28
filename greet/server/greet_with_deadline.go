package main

import (
	"context"
	"log"
	"time"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GreetWithDeadline(ctx context.Context, in  *pb.GreetRequest) (*pb.GreetResponse, error){
	log.Printf("GreetWithDeadline function was invoked with %v", in)
	for i := 0; i <= 3 ; i++{
		if ctx.Err() == context.DeadlineExceeded{
			log.Println("the client cancled the request!!")
			return nil, status.Error(codes.Canceled, "the client cancle the request!!")
		}
		time.Sleep(1*time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello with Deadline, " + in.Name,
	},nil
}