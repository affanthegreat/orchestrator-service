package main

import (
	"context"
	"errors"
	"log"
	"net"
	p "orca2"
	q "orca3"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	p.UnimplementedStatusUpdaterServer
}

func main() {
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()

	p.RegisterStatusUpdaterServer(srv, &server{})
	log.Printf("server listening at %xv", listener.Addr())
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) GetUser(ctx context.Context, name *p.ValidUserName) (*p.StatusTwo, error) {
	if name.Statuscode == "ok" {
		conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		client := q.NewAddServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		req := &q.SearchQuery{Query: name.Query}
		r, err := client.GetMockUserData(ctx, req)
		if err != nil {
			log.Fatalf("Error at second Orcestration: %v", err)
		}
		if r != nil {
			log.Printf("Connection sucessful to second orchestrator")
			log.Println(r.Class)
			log.Println(r.Name)
			log.Println(r.Roll)
			return &p.StatusTwo{StatusCode: 200, Name: r.Name, Class: r.Class, Roll: r.Roll}, nil
		}
		return nil, errors.New("r is empty")
	}
	return nil, errors.New("Can't call directly to second orchestrator directly.")
}
