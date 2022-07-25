package main

import (
	"context"
	"log"
	"time"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreeet function was invoked")

	stream , err := c.LongGreet(context.Background())
	if err != nil{
		log.Fatalf("error while calling LongGreet : %v", err)
	}
	reqs := []*pb.GreetRequest{
		{Name: "Ade"},
		{Name: "Abdi"},
		{Name: "Henok"},
	}
	for _, req := range reqs{
		log.Printf("request : %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err  := stream.CloseAndRecv()

	if err != nil{
		log.Printf("error while return response with : %v",err)
	}

	log.Printf("request responses \n", res.Result)
}