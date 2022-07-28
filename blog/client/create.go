package main

import (
	"context"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
)

func doCreate(c pb.BlogServiceClient)string {
	log.Println("___doCreate fucn was invoked___")
	blg := &pb.Blog{
		AuthorId: "abdi",
		Title: "life",
		Content: "life is fun ",
	}
	res, err := c.CreateBlog(context.Background(), blg)
	if err != nil{
		log.Fatalf("error while requesting create method in client %v", err)
	}
	log.Printf("The id is : %v", res.Id)
	return res.Id
}