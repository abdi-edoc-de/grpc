package main

import (
	"log"
	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var adrs = "localhost:50051"
func main() {
	conn , err := grpc.Dial(adrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("error wihle creating connection in client file : %v", err)
	}
	c:=pb.NewBlogServiceClient(conn)
	// doRead(c, doCreate(c))
	// blog := doRead(c, doCreate(c))
	// blog.Title = "life updated"
	// doUpdate(c, blog)
	// doListBlog(c)
	id := doCreate(c)
	doRead(c, id)
	doDelelte(c, id)
	doRead(c, id)
	



}
