package main
import (
	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
	"io"
	"log"
)

func (s *server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("Average function was invoked : %v", stream)
	value, nums := int32(0), int32(0)
	for {
		req, err := stream.Recv()
		if err ==  io.EOF{
			return stream.SendAndClose(&pb.AverageResponse{
				Result : value/nums,
			})
		}
		value += req.Number
		nums++
	}
	return nil
}