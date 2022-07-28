package main

import (
	"context"
	"io"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func doListBlog(c pb.BlogServiceClient) []*pb.Blog{
	log.Println("doListBlog function was invoked")
	blogs := []*pb.Blog{}
	stream, err := c.ListBlog(context.Background() , &emptypb.Empty{})
	if err != nil{
		log.Fatalf("error while creating the stream : %v", err)
	}
	for {
		res,err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("error while reading from stream in doListBlog func with error : %v", err)
		}
		log.Printf("From list of blogs : %v" , res)
		blogs = append(blogs, res)
	}
	return blogs
}