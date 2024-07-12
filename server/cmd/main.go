package main

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
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

func (h *handlerServer) SendBigData(ctx context.Context, req *pb_impl.SendNotificationBatchRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, v := range md {
			fmt.Printf(" %s: %s\n", k, strings.Join(v, ","))
		}
	}
	fmt.Printf("Received data: Data Length=%d bytes\nContent: %+v\n\n", len(req.GetPersonalizations()), req.GetPersonalizations())
	fmt.Printf("Received data: Text Length=%d bytes\nContent: %+v\n\n", len(req.GetText()), req.GetText())

	return &emptypb.Empty{}, nil
}

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
	var data bytes.Buffer
	for {
		req, err := sv.Recv()
		if err == io.EOF {
			if err := sv.SendAndClose(&pb_impl.DataResponse{Message: "file uploaded!"}); err != nil {
				log.Println(err)
				return err
			}
			break
		}
		if err != nil {
			return err
		}
		chunk := req.GetData()
		_, err = data.Write(chunk)
		if err != nil {
			log.Println(err)
		}
	}

	f, err := os.Create("end.pdf")
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = data.WriteTo(f)

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
