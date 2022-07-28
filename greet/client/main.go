package main

import (
	"log"
	"time"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var address = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls{
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		
		if err != nil{
			log.Fatalf("error while loading certificates error: %v",err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

	}
	conc, err := grpc.Dial(address,opts...) 
	
	if err != nil{
		log.Fatalf("failed to connect: %v", err)
	}
	defer conc.Close()

	c := pb.NewGreetServiceClient(conc)
	doGreetWithDeadline(c,  5* time.Second)
}