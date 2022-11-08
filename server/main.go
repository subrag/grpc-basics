package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net"

	pb "github.com/subrag/grpc-basics/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:5002"

type Server struct {
	pb.ProjectServiceServer
	dbProj       []*pb.Project
	dbUser       []*pb.UserProfile
	dbAssignment []*pb.Assignment
}

func main() {
	user, proj, asgmt := LoadData()
	log.Print(user, proj, asgmt)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", addr)
	}
	defer lis.Close()

	tls := false
	opts := []grpc.ServerOption{}

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewClientTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Error in TLS: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterProjectServiceServer(s, &Server{
		dbUser:       user,
		dbProj:       proj,
		dbAssignment: asgmt,
	})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Fail to serve")
	}

}

func LoadData() ([]*pb.UserProfile, []*pb.Project, []*pb.Assignment) {
	var data []byte
	// load users data
	data, err := ioutil.ReadFile("test_data/users.json")
	if err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	var up []*pb.UserProfile
	if err := json.Unmarshal(data, &up); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	log.Println(up)

	// load project data
	data, err = ioutil.ReadFile("test_data/project.json")
	if err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	var proj []*pb.Project
	if err := json.Unmarshal(data, &proj); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
	log.Println(proj)

	// create assignment out of project and users
	var asgmt []*pb.Assignment
	n := 0
	for n < len(proj) {
		as := []*pb.UserProfile{
			up[rand.Intn(len(up))],
			up[rand.Intn(len(up))],
		}
		asgmt = append(asgmt, &pb.Assignment{
			Project:  proj[n],
			Assignee: as,
		})

		n += 1
	}
	log.Println(asgmt)
	return up, proj, asgmt
}
