package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/ervitis/grpc-test-unary-exhausted/pb_impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	f, err := os.Open("../../test_files/biggrpc.pdf")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()
	
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial("localhost:8490", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb_impl.NewDataServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := c.SendData(ctx, &pb_impl.DataRequest{Data: data})
	if err != nil {
		panic(err)
	}
	log.Println(r.GetMessage())
}
