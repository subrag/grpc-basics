package main

import (
	"context"
	"io"
	"log"
	"time"

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
			log.Println("Done reading stream!.")
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v.", err)
		}

		log.Printf("Id: %v Project: %v\n", msg.GetId(), msg.Name)
	}

}

func createProjects(c pb.ProjectServiceClient) ([]*pb.Project, error) {
	newProjs := []*pb.Project{
		{Name: "SODEXO", Id: 2001},
		{Name: "PHARMEasy", Id: 2002},
		{Name: "LSOS", Id: 2003},
	}

	stream, err := c.CreateProjects(context.Background())
	if err != nil {
		log.Fatalf("Create project failed with error: %v", err)
	}

	for _, newProj := range newProjs {
		log.Printf("Sending project: %v", newProj.Name)
		stream.Send(newProj)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from create projects %v\n", err)
	}
	log.Printf("Projects created: %v.", res.Msg)
	return newProjs, nil
}

func createProjAssignment(c pb.ProjectServiceClient, projs []*pb.Project) {
	log.Println("Create project assignment started.")
	stream, err := c.CreateAssignments(context.Background())
	if err != nil {
		log.Fatalf("Error while intiating client streaming: %v.", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Falied to receive stream. Error: %v\n", err)
				break
			}
			log.Printf("Created assignment: %v", res)
		}
		// close(waitc)
	}()

	go func() {
		for _, proj := range projs {
			log.Printf("Create Assignment for %s", proj.GetName())
			stream.Send(proj)
			time.Sleep(1 * time.Second)
		}
	}()

	stream.CloseSend()
	<-waitc
}
