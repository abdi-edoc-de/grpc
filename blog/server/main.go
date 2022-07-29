package main

import (
	"context"
	"log"
	"net"

	pb "github.com/abdi-edoc-de/grp_course_pro/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var adrs = "0.0.0.0:50051"
type server struct{
	pb.BlogServiceServer
}
var collection *mongo.Collection

func main(){
	client , err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil{
		log.Fatalf("error while creating connection to mongodb with error: %v", err)
	}
	err1 := client.Connect(context.Background())
	if err1 != nil{
		log.Fatalf("error while connecting to db with error %v", err)
	}
	collection = client.Database("BlogDb").Collection("Blog")
	lis, err := net.Listen("tcp", adrs)
	if err != nil{
		log.Fatalf("error while trying to listen on %v address  with error %v",adrs, err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(unaryIntercept))
	pb.RegisterBlogServiceServer(s, &server{})
	log.Printf("the server is listening on %v", adrs)
	if err :=  s.Serve(lis); err != nil{
		log.Fatalf("error while listing with error : %v", err)
	}
}
