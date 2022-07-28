package main

import (
	"context"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
	"google.golang.org/grpc/status"
)

func doDelelte(c pb.BlogServiceClient, id string) {
	log.Printf("doDelelte function was invoked with id")
	_, err := c.DeleltBlog(context.Background(), &pb.BlogId{
		Id: id,
	})
	if err != nil{
		e , ok := status.FromError(err)
		if ok{
			log.Fatalf("grpc errors : %v", e)
		}
		log.Fatalf("error while calling the server method called delelteBlog with error : %v", err)
	}
	log.Println("The delelte task has succeded")
}
