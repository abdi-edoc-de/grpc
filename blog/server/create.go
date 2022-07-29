package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error){
	log.Printf("CreateBlog funtion was invoked with %v", in)
	blg := BlogItem{
		AuthorId: in.AuthorId,
		Title: in.Title,
		Content: in.Content,
	}
	res , err := collection.InsertOne(ctx, blg)
	if err != nil{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal errror : %v",err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error : %v\n",err),
		)
	}
	time.Sleep(10 * time.Second)

	return &pb.BlogId{
		Id : oid.Hex(),
	}, nil

}
