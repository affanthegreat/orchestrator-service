package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	p "part2/proto"

	"google.golang.org/grpc"
)

type server struct {
	p.UnimplementedAddServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()

	p.RegisterAddServiceServer(srv, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
func (s *server) GetUserByName(ctx context.Context, name *p.SearchQuery) (*p.User, error) {
	a := name.Query
	fmt.Print(a)
	return nil, errors.New("not implemented yet." + name.Query + " will implement me")
}
