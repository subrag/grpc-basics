package main

import (
	"context"
	"time"

	"log"

	pb "github.com/subrag/grpc-basics/proto"
)

func (s *Server) GetAssignment(ctx context.Context, p *pb.Project) (*pb.Assignment, error) {
	log.Printf("Get assignment called...")
	return s.dbAssignment[0], nil

}

func (s *Server) GetAllProjects(in *pb.EmptyRequest, stream pb.ProjectService_GetAllProjectsServer) error {
	log.Print("Server streaming")
	for i := 0; i < len(s.dbProj); i++ {
		time.Sleep(2 * time.Second)
		stream.Send(s.dbProj[i])
	}
	return nil
}
