package main

import (
	"context"
	"io"
	"log"

	pb "github.com/subrag/grpc-basics/proto"
)

func getAssignment(c pb.ProjectServiceClient) {
	res, err := c.GetAssignment(context.Background(), &pb.Project{
		Name: "ABC",
		Id:   12,
	})
	if err != nil {
		log.Fatalf("Faled to fetch assignments, with error: %v", err)
	}
	log.Printf("Retunred assignments: %v", res.Assignee)

}

func getAllProjects(c pb.ProjectServiceClient) {
	log.Println("Get all project invoked")

	req := &pb.EmptyRequest{}
	stream, err := c.GetAllProjects(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to fetch all projetcs with error: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("Reached EOF")
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v.", err)
		}

		log.Printf("Id: %v Project: %v\n", msg.GetId(), msg.Name)
	}

}
