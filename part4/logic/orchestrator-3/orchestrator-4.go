package main

import (
	"context"
	"errors"
	"log"
	"net"
	p "orca3"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	p.UnimplementedAddServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()

	p.RegisterAddServiceServer(srv, &server{})
	log.Printf("server listening at %xv", listener.Addr())
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) GetMockUserData(ctx context.Context, name *p.SearchQuery) (*p.User, error) {
	a := name.Query
	b := len([]rune(a))
	c, err := strconv.ParseInt(strconv.Itoa(b*10), 10, 64)
	if b < 6 {
		return nil, errors.New("Length of Name smaller than 6 characters")
	} else if err == nil {
		log.Println(a)
		log.Println(strconv.Itoa(b))
		log.Println(c)
		return &p.User{Name: a, Class: strconv.Itoa(b), Roll: c}, nil
	}
	return nil, err
}
