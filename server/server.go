package server

import (
	"fmt"
	"github.com/just1689/fun-with-grpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCommServer(s, server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

type server struct{}

func (s server) Ping(srv proto.Comm_PingServer) error {

	log.Println("start new server")
	ctx := srv.Context()

	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("Clean exit")
			return nil
		}
		if err != nil {
			log.Printf("GRPC stream err %v", err)
			continue
		}

		fmt.Println("Received: ", req.Val)

		resp := proto.Msg{Val: "Pong"}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
	}
}
