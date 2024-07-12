package main

import (
	"context"
	"fmt"
	"github.com/ervitis/grpc-test-unary-exhausted/client/data"
	"github.com/ervitis/grpc-test-unary-exhausted/pb_impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func chunkItems[T any](items []*T, chunkSize int, fn func(chunk []*T, loopCount int) error) error {
	for i := 0; i < len(items); i += chunkSize {
		end := i + chunkSize
		if end > len(items) {
			end = len(items)
		}
		chunk := items[i:end]
		err := fn(chunk, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	/*
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
	conn, err := grpc.Dial("localhost:8590",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
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

	/*
		// now try to send it on stream
		stream, err := c.SendStream(ctx)
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(f)

		// modify size
		const size = 8192
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
	*/

	msgs := data.NewRpcMessage(80000)
	chunkSize := 500
	if err := chunkItems(msgs, chunkSize, func(chunk []*pb_impl.Personalization, loopCount int) error {
		return sendChat(ctx, c, msgs, data.MockMessage(10000))
	}); err != nil {
		panic(err)
	}
	for i := 0; i < len(msgs); i += chunkSize {
		d := msgs[i : i+chunkSize]
		_, err = c.SendBigData(ctx, &pb_impl.SendNotificationBatchRequest{
			Personalizations: d,
			Text:             data.MockMessage(100000),
		})
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("End!")
}

func sendChat(ctx context.Context, c pb_impl.DataServiceClient, personalizations []*pb_impl.Personalization, text string) error {
	_, err := c.SendBigData(ctx, &pb_impl.SendNotificationBatchRequest{
		Personalizations: personalizations,
		Text:             text,
	})
	if err != nil {
		return err
	}
	return nil
}
