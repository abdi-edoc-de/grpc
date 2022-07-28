package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient){
	log.Printf("doMax funtion was invoked with : %v", c)
	stream , err := c.Max(context.Background())
	if err != nil{
		log.Fatalf("error while creating stream in client file with : %v", err)
	}
	waitc := make(chan struct{})
	go func ()  {
		reqs := []*pb.MaxRequest{
			{Number: 1},
			{Number: 0},
			{Number: 4},
			{Number: 10},
			{Number: 2},
			{Number: 100},
		}
		for _, req := range reqs{
			log.Printf("sending request : %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func ()  {
		for {
			res, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Fatalf("error while reading from stream in client file : %v", err)
			}
			log.Printf("The current max is : %v", res.Result)
		}	
		close(waitc)
	}()
	<- waitc
}