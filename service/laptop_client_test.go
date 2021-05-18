package service

import (
	"context"
	"net"
	"testing"

	"github.com/samirprakash/go-grpc-pc-book/pb"
	"github.com/samirprakash/go-grpc-pc-book/sample"
	"github.com/samirprakash/go-grpc-pc-book/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestCreateLaptopClient(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// check that the laptop is saved in the store
	other, err := laptopServer.laptopStore.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check if the saved laptop is the same which was created
	jsonLaptop, err := serializer.ProtobufToJSON(laptop)
	require.NoError(t, err)

	jsonOther, err := serializer.ProtobufToJSON(other)
	require.NoError(t, err)

	require.Equal(t, jsonLaptop, jsonOther)
}

func startTestLaptopServer(t *testing.T) (*LaptopServer, string) {
	laptopServer := NewLaptopServer(NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listner, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listner)

	return laptopServer, listner.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}
