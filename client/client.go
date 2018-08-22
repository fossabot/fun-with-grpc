package main

import (
	"context"
	"github.com/just1689/fun-with-grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connecting to grpc endpoint failed with err: %v", err)
	}

	client := proto.NewCommClient(conn)
	stream, err := client.Ping(context.Background())
	if err != nil {
		log.Fatalf("Failed to open GRPC stream: %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)

	go func() {

		req := proto.Msg{Val: "This message is from the client!"}
		if err := stream.Send(&req); err != nil {
			log.Fatalf("Failed to send: %v", err)
		}
		log.Printf("Sending: %s", req.Val)
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}

	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive from server %v", err)
			}
			log.Printf("Recieved from the server: %s", resp.Val)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()
	<-done

}
