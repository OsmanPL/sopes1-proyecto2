package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"

	pb "github.com/sevenpok/api-rabbit/gen/proto"

	"google.golang.org/grpc"
)
const(
	topic = "Game"
	brokerAddress = "localhost:9092"
)
func produce (ctx context.Context, req *pb.Game){
	l := log.New(os.Stdout, "Kafka Escribiendo: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic: topic,
		Logger: l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(int(req.Id))),
		Value: []byte("{id: "+ strconv.Itoa(int(req.Id))+ ", gameName: "+req.GameName+", players: "+strconv.Itoa(int(req.Players))+"}"),
	})

	if err != nil {
		panic("No se pudo enviar el mensaje: "+ err.Error())
	}
}
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
