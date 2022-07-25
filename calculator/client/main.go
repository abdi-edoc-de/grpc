package main

import (
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


var address = "localhost:50051"
func main(){
	conn , err := grpc.Dial(address,  grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Fatalf("failed to connect : %v", err)
	}
	
	c := pb.NewCalculatorServiceClient(conn)
	doAverage(c)
}