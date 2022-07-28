package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Printf("ListBlog function was invoked with %v",in)
	cur , err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil{
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Uknown internal error"),

		)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()){
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil{
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("error while decoding the data error: %v", err),
			)
		}
		stream.Send(documentToBlog(data))

	}
	if err := cur.Err(); err != nil{
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error : %v", err),
		)
	}

	return nil
}