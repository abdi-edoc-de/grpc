package main

import (
	"context"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
)
func doSum(c pb.CalculatorServiceClient){
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		X:10, Y:20,
	})
	if err != nil{
		log.Fatalf("can not call sum rpc function : %v", err)
	}
	log.Printf("The sum is : %v", res.Result)
}