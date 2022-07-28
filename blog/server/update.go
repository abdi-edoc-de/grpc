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

func (s *server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("---UpdateBlog function was invoked with %v---", in)
	oid , err := primitive.ObjectIDFromHex(in.Id)
	if err != nil{
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Can not parse Id %v", err),
		)
	}
	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:  in.Title,
		Content: in.Content,
	}
	
	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id":oid},
		bson.M{"$set":data},
	)
	if err != nil{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error with %v", err),
		)
	}
	
	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Can not A blog with Id %v", in.Id),
		)
	}
	log.Println("___updated___")
	return &emptypb.Empty{}, nil
}