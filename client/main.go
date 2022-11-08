package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/subrag/grpc-basics/proto"
)

var addr string = "0.0.0.0:5002"

func main() {
	tls := false

	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error tls: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v, Error: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewProjectServiceClient(conn)
	// getAssignment(c)

	getAllProjects(c)

	// projs, err := createProjects(c)
	// if err != nil {
	// 	log.Fatalf("Error while creating projects: %v.", err)
	// }
	// createProjAssignment(c, projs)

}
