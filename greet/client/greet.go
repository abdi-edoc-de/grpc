package main

import (
	"context"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/greet/proto"
)

func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		Name: "Abdi",
	})
	if err != nil{
		log.Fatal("greet got to error :%v", err)
	}
	log.Printf("Greeting: %v", res.Result)
}