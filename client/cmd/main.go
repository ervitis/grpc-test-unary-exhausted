package main

import (
	"context"
	"log"

	"github.com/ervitis/grpc-test-unary-exhausted/pb_impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8490", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb_impl.NewDataServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := c.SendData(ctx, &pb_impl.DataRequest{Data: []byte(`hello`)})
	if err != nil {
		panic(err)
	}
	log.Println(r.GetMessage())
}
