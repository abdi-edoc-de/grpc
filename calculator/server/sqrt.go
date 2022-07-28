package main

import (
	"context"
	"fmt"
	"math"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	number := in.Number
	if number < 0{
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("the argument number can not be negative : %v", number),
		)
	}
	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}