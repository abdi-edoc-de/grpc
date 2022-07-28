package main

import (
	"log"
	"context"
	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
)

func doUpdate(c pb.BlogServiceClient, updatedBlog *pb.Blog) {
	log.Printf("doUpdate function was invoked with %v", updatedBlog)
	_, err := c.UpdateBlog(context.Background(), updatedBlog)
	if err != nil{
		log.Fatalf("error while trying to update in clinet part with error %v" , err)
	}
	log.Println("the blog has been updated")
}
