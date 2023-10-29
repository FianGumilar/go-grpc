package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"sync"

	pb "github.com/FianGumilar/go-grpc/student"
	"google.golang.org/grpc"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex // if happens race condition use mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {

	log.Println("Incoming request for Student By Email")

	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	data, err := os.ReadFile("data/data.json")
	if err != nil {
		log.Printf("error read data file: %v", err.Error())
	}

	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalf("error unmarshal data: %v", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalf("error listening: %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("error when serve grpc: %v", err.Error())
	}
}
