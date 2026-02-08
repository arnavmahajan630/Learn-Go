package main

import (
	"log"
	"net"
	pb "github.com/arnavmahajan630/Learn-Go/Intermeidate-Projects/go-grpc/proto"
	"google.golang.org/grpc"
)

// create the server for grpc requests

type helloServer struct {
	pb.GreetServiceServer 
}

func main() {
	 
	l , err := net.Listen("tcp", ":8080");
	if err != nil {
		log.Fatal("Failed to start the server")
	}
	serve := grpc.NewServer()
	
	pb.RegisterGreetServiceServer(serve, &helloServer{})
	log.Printf("Server started on Port 8080")

	if err := serve.Serve(l); err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}


}
