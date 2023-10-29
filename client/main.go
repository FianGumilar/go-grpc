package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/FianGumilar/go-grpc/student"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalf("error in dialing: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "fian@gmail.com")
	getDataStudentByEmail(client, "vania@gmail.com")
}

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Printf("error get student by email: %v", err.Error())
	}

	fmt.Println(student)

}
