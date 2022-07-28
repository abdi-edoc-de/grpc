package main

import (
	// "fmt"
	"io"
	"log"
	"math"

	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
)

func (s *server) Max(stream pb.CalculatorService_MaxServer) error {
	current := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			log.Fatalf("error while reciving %v", err)
		}
		current = math.Max(float64(current), float64(req.Number))
		if  err := stream.Send(&pb.MaxResponse{Result: int32(current)}); err != nil{
			log.Fatalf("error while reading from stream in server", err)
		}

	}
	return nil
}