package main

import (
	"context"
	"log"
	"time"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration){
	log.Println("doGreetWithDeadline function was invoked")
	ctx , cancle := context.WithTimeout(context.Background(), timeout)
	defer cancle()
	req := &pb.GreetRequest{
		Name: "Abdi",
	}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil{
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded{
				log.Printf("deadline exceeded Message : %v with code : %v", e.Message(), e.Code())
				return
			}
			log.Printf("another gRpc errror happend error: %v , code : %v", e.Message(), e.Code())
			return
		}
		log.Fatalf("error while requesting response in client file with error: %v", err)
	}
	log.Printf("GreetingWithDeadling : %v", res.Result)
}

