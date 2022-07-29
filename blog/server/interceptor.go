package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)
func unaryIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(interface{}, error){
	fmt.Println("___----_intercepted----____hi___----")
	hnd, err := handler(ctx, req)
	if err != nil{
		log.Fatalf("gRpc error happend with error : %v", err)
	}
	return hnd, nil
}