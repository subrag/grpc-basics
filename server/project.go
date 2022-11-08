package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"

	"log"

	pb "github.com/subrag/grpc-basics/proto"
)

func (s *Server) GetAssignment(ctx context.Context, p *pb.Project) (*pb.Assignment, error) {
	log.Printf("Get assignment called...")
	return s.dbAssignment[0], nil

}

func (s *Server) GetAllProjects(in *pb.EmptyRequest, stream pb.ProjectService_GetAllProjectsServer) error {
	log.Print("Server streaming.")
	for i := 0; i < len(s.dbProj); i++ {
		time.Sleep(2 * time.Second)
		stream.Send(s.dbProj[i])
	}
	return nil
}

func (s *Server) CreateProjects(stream pb.ProjectService_CreateProjectsServer) error {
	log.Println("CreateProjects function invoked.")
	res := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.ResponseStatus{Msg: res})
		}
		if err != nil {
			log.Fatalf("Error recieving stream.")
		}
		s.dbProj = append(s.dbProj, req)
		res = fmt.Sprintf("%s %s", res, req.GetName())
	}
}

func (s *Server) CreateAssignments(stream pb.ProjectService_CreateAssignmentsServer) error {
	log.Println("Bidirectional streaming")

	for {
		proj, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while sending to client, error: %v\n", err)
		}
		asgmt, err := createProjAssgmtLogic(proj, s.dbUser)
		if err != nil {
			log.Fatalf("Faled to create project assignment for %s, with error: %v", proj.Name, err)
		}
		stream.Send(asgmt)

	}

}

func createProjAssgmtLogic(p *pb.Project, allUsers []*pb.UserProfile) (*pb.Assignment, error) {
	n := rand.Intn(len(allUsers))
	var users []*pb.UserProfile
	// Making team size random by selcting n random
	for i := 0; i < n; i += 1 {
		users = append(users, allUsers[rand.Intn(len(allUsers))])
	}
	log.Printf("Adding assignment: %v - %v", p, users)
	time.Sleep(3 * time.Second)
	return &pb.Assignment{
		Project:  p,
		Assignee: users,
	}, nil
}
