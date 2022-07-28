package main

import (
	"context"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
	"google.golang.org/grpc/status"
)

func doRead(c pb.BlogServiceClient, id string) *pb.Blog{
	log.Printf("doRead function was invoked with id : %v", id)
	res, err := c.ReadBlog(context.Background(), &pb.BlogId{
		Id: id,
	})
	if err != nil{
		e , ok := status.FromError(err)
		if ok{
			log.Fatalf("grpc errors : %v", e)
		}
		log.Fatalf("error while trying read from the server %v", err)
	}
	log.Printf("The blog that has been red: %v", res)
	return res
}