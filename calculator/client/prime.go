package main

import (
	"context"
	"io"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient){
	stream , err := c.Prime(context.Background(), &pb.PrimeRequest{
		Number: 1050,
	})	
	if err != nil{
		log.Fatalf("error while calling the prime function: %v", err)
	}
	for {
		msg, err := stream.Recv();
		if err == io.EOF{break}
		if err != nil{
			log.Fatalf("error while reading from the stream : %v", err)
		}

		log.Printf("the factor: %v", msg.Factor)
	}
}