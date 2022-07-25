package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)

var address = "localhost:50051"

func main() {
	conc, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials())) 
	
	if err != nil{
		log.Fatalf("failed to connect: %v", err)

	}
	defer conc.Close()
	c := pb.NewGreetServiceClient(conc)
	doLongGreet(c)
}