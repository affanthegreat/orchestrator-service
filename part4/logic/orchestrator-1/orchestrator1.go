package main

import (
	"context"
	"errors"
	"log"
	"net"
	p "orca1"
	q "orca2"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	p.UnimplementedAddServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
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

var s = []string{"James", "Wagner", "Christene", "Mike"}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func checkUserName(name string) bool {
	return contains(s, name)
}

func (s *server) GetUserByName(ctx context.Context, name *p.SearchQuery) (*p.Status_One, error) {

	if checkUserName(name.Query) {
		//connect to second orchestrator
		conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client := q.NewStatusUpdaterClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.GetUser(ctx, &q.ValidUserName{Query: name.Query, Statuscode: "ok"})
		if err != nil {
			log.Printf("Error at second Orcestration: %v", err)
			return nil, err
		}
		log.Printf("Connection sucessful to second orchestrator")
		if r != nil {

			return &p.Status_One{Code: 200, Name: r.Name, Class: r.Class, Roll: r.Roll}, nil
		}
	} else {
		return nil, errors.New("User does not exist in given data. ")
	}
	return nil, errors.New("User does not exist in given data. ")
}
