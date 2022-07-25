package main

import (
	"context"
	"log"
	"time"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
)


func doAverage(c pb.CalculatorServiceClient){
	log.Println("doAverage function was invoked")
	reqs := []*pb.AverageRequest{
		{Number: 3},
		{Number: 5},
		{Number: 8},
		{Number: 10},
	}
	stream , err := c.Average(context.Background())
	if err != nil{
		log.Fatalf("error while calling average : %v", err)
	}
	for _, req := range reqs {
		log.Printf("send ing value : %v \n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res , err:= stream.CloseAndRecv()
	if err != nil{
		log.Printf("error while return response in client : %v", err)
	}
	log.Printf("The Aerage of %v is : %v", reqs, res.Result)
}