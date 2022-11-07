package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/subrag/grpc-basics/proto"
)

var addr string = "0.0.0.0:5002"

func main() {
	// r := gin.Default()

	// r.GET("/status", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"status": "OK",
	// 	})
	// })
	// r.Run(":5001")

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v, Error: %v", addr, err)
	}
	defer conn.Close()

	c := pb.NewProjectServiceClient(conn)
	// getAssignment(c)

	// getAllProjects(c)

	projs, err := createProjects(c)
	if err != nil {
		log.Fatalf("Error while creating projects: %v.", err)
	}
	createProjAssignment(c, projs)

}
