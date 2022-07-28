package main

import (
	"context"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient){
	log.Println("doSqrt function was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: -9})
	if err != nil{
		e , ok := status.FromError(err)
		if ok{
			log.Printf("error message from server :%v", e.Message())
			log.Printf("error code from server : %v", e.Code())
			if e.Code() == codes.InvalidArgument{
				log.Printf("you cant put negative number as argument")
			}
			log.Fatalf("error after checking")
		}else{
			log.Fatalf("error while requesting from client in client file : %v", err)
		}
	}
	log.Printf("The square root is : %v", res.Result)
}