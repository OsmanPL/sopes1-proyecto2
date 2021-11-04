package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sevenpok/api-rabbit/gen/proto"
	"github.com/streadway/amqp"

	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func producerRabbit(req *pb.Game) {
	msg := fmt.Sprintf("{id:%d,gameName:%s,players:%d}", req.Id, req.GameName, req.Players)

	rabbit, err := amqp.Dial("amqp://admin:admin@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()

	channel, err := rabbit.Channel()

	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare("game", false, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}
	err = channel.Publish("", queue.Name, false, false,
		amqp.Publishing{
			Headers:     nil,
			ContentType: "text/plain",
			Body:        []byte(msg),
		})

	if err != nil {
		log.Fatal(err)
	}
}

func (s *testApiServer) CreateGame(ctx context.Context, req *pb.Game) (*pb.ResponseRequest, error) {
	fmt.Println(req)
	producerRabbit(req)
	msg := pb.ResponseRequest{Msg: "creado"}
	return &msg, nil
}

func main() {
	//Escuchamos al servidor grpc
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	log.Println("Server Listening on port 8080")
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Panicln((err))
	}
}
