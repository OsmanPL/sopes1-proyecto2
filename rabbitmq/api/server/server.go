package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sevenpok/api-rabbit/gen/proto"

	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) CreateGame(ctx context.Context, req *pb.Game) (*pb.ResponseRequest, error) {
	fmt.Println(req)
	msg := pb.ResponseRequest{Msg: "creado"}
	return &msg, nil
}

func main() {
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Panicln((err))
	}
}
