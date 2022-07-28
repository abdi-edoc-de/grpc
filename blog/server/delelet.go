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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) DeleltBlog(ctx context.Context,in  *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("Delelet Blog was invoked with %v", in)
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cant parse Id error: %v", err),
		)
	}
	res , err := collection.DeleteOne(ctx, bson.M{"_id":oid})

	if err != nil{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error with %v", err),
		)
	}
	if res.DeletedCount == 0{
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cant find with id : %v", in.Id),
		)
	}

	return &emptypb.Empty{}, nil
}