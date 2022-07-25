package main

import (
	"context"
	"io"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient){
	log.Printf("doGreetManyTimes Funtion was invoked")
	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		Name: "abdi",
	})
	if err != nil{
		log.Fatalf("error while calling greetManyTimes rp : %v",err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("error while reading the stream : %v", err)
		}
		log.Printf("Greeting : %v", msg.Result)
	}
}