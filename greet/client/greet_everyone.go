package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone function was invoked")
	stream , err := c.GreetEveryone(context.Background())
	if err != nil{
		log.Fatalf("error while creating stream in client : %v", err)
	}
	waitc := make(chan struct{})
	go func ()  {
		reqs := []*pb.GreetRequest{
			{Name: "Ade"},
			{Name: "Abdi"},
			{Name: "Henok"},
		}
		for _, req := range reqs{
			log.Printf("sending req to server : %v", req)
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
				log.Fatalf("error while reading from server in client file : %v", err)
			}
			log.Printf("Greeting from server: %v\n", res.Result)
		}
		close(waitc)
	}()

	<- waitc
	
}