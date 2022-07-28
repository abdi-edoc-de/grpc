package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// primitive.ObjectIDFromHex(id)

// fileter := bson.M("_id":oid)

// res.Decode(data)
func (s *server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error){
	log.Printf("ReadBlog function was invoked with %v", in)
	oid , err := primitive.ObjectIDFromHex(in.Id)
	if err != nil{
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("invalid id as argument"),
		)
	}
	filter := bson.M{"_id":oid}
	res := collection.FindOne(ctx, filter)
	data := &BlogItem{}
	if err := res.Decode(data); err != nil{
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can not found the blog with Id %v", in.Id),
		)
	}
	return documentToBlog(data), nil
}