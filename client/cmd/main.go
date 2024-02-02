package main

import (
	"bufio"
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

	/*
		data, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
	*/
	conn, err := grpc.Dial("localhost:8590", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb_impl.NewDataServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	/*
		r, err := c.SendData(ctx, &pb_impl.DataRequest{Data: data})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(r.GetMessage())
		}
	*/

	// now try to send it on stream
	stream, err := c.SendStream(ctx)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)

	// modify size
	const size = 4096
	buff := make([]byte, size)

	for {
		n, err := reader.Read(buff)
		if err == io.EOF {
			break
		}
		if err := stream.Send(&pb_impl.DataRequest{Data: buff[:n]}); err != nil {
			panic(err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	log.Println(resp.GetMessage())
}
