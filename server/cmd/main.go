package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ervitis/grpc-test-unary-exhausted/pb_impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type handlerServer struct{}

func (h *handlerServer) SendData(ctx context.Context, req *pb_impl.DataRequest) (*pb_impl.DataResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, v := range md {
			fmt.Printf(" %s: %s\n", k, strings.Join(v, ","))
		}
	}
	fmt.Printf("Received data: Data Length=%d bytes\n", len(req.Data))

	return &pb_impl.DataResponse{Message: "hello"}, nil
}

func (h *handlerServer) SendStream(sv pb_impl.DataService_SendStreamServer) error {
	return nil
}

func main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	lis, err := net.Listen("tcp", ":8590")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb_impl.RegisterDataServiceServer(server, &handlerServer{})

	go func() {
		log.Panic(server.Serve(lis))
	}()

	<-s
	server.GracefulStop()
	log.Println("stopped")
}
