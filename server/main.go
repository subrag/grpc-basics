package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net"

	pb "github.com/subrag/grpc-basics/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:5002"

type Server struct {
	pb.ProjectServiceServer
	dbProj       []*pb.Project
	dbUser       []*pb.UserProfile
	dbAssignment []*pb.Assignment
}

func main() {
	// r := gin.Default()

	// r.GET("/status", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"status": "OK",
	// 	})
	// })
	// r.Run(":5002")
	user, proj, asgmt := LoadData()
	log.Print(user, proj, asgmt)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", addr)
	}
	s := grpc.NewServer()
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
