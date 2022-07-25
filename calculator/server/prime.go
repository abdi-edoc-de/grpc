package main

import (
	"log"
	pb "github.com/abdi-edoc-de/grp_course_pro/calculator/proto"
)

func (s *server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error{
	log.Printf("Prime function was invoked : %v", in)
	N, k := in.Number, int32(2)
	for N > 1{
		if (N % k == 0){   // if k evenly divides into N
        	// print k      // this is a factor
			stream.Send(&pb.PrimeResponse{
				Factor: k,
			})
        	N = N / k    // divide N by k so that we have the rest of the number left.
		}else{	
			k = k + 1
		}
	}
	return nil
}
