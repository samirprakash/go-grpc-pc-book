package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/samirprakash/go-grpc-pc-book/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	laptopStore LaptopStore
}

// NewLaptopServer returns a new laptop server
func NewLaptopServer(laptopStore LaptopStore) *LaptopServer {
	return &LaptopServer{laptopStore: laptopStore}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("received a create laptop request with id : %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// check if the user provided Id is a valid UUID or not
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID : %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate laptp id : %v", err)
		}
		laptop.Id = id.String()
	}

	time.Sleep(6 * time.Second)

	if ctx.Err() == context.Canceled {
		log.Print("request canceled by client")
		return nil, status.Error(codes.Canceled, "canceled by client")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	// save the laptop to in memory store
	err := server.laptopStore.Save(laptop)

	// check if error is because of an existing laptop in the meory store
	// or due to an internal server error
	errCode := codes.Internal
	if errors.Is(err, ErrAlreadyExists) {
		errCode = codes.AlreadyExists
	}

	if err != nil {
		return nil, status.Errorf(errCode, "not able to save laptop to in memery store : %v", err)
	}

	log.Printf("saved laptop with Id : %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}
