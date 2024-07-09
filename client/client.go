package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "swag-grpc-crud/proto"
)

func main() {
	// Establish a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client for UserService
	c := pb.NewUserServiceClient(conn)

	// Create users
	users := []struct {
		name  string
		email string
	}{
		{"John Doe", "john.doe@example.com"},
		{"Jane Smith", "jane.smith@example.com"},
		{"Alice Johnson", "alice.johnson@example.com"},
	}

	var createdUsers []*pb.CreateUserResponse
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Loop to create users
	for _, user := range users {
		// Make gRPC CreateUser request
		createRes, err := c.CreateUser(ctx, &pb.CreateUserRequest{Name: user.name, Email: user.email})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf("User created: %v", createRes)
		createdUsers = append(createdUsers, createRes)
	}

	// Update a user
	updateReq := &pb.UpdateUserRequest{
		Id:    createdUsers[0].GetId(),
		Name:  "John Updated",
		Email: "john.updated@example.com",
	}
	updateRes, err := c.UpdateUser(ctx, updateReq)
	if err != nil {
		log.Fatalf("could not update user: %v", err)
	}
	log.Printf("User updated: %v", updateRes)

	// Get users
	for _, createdUser := range createdUsers {
		// Make gRPC GetUser request
		getRes, err := c.GetUser(ctx, &pb.GetUserRequest{Id: createdUser.GetId()})
		if err != nil {
			log.Fatalf("could not get user: %v", err)
		}
		log.Printf("User fetched: %v", getRes)
	}

	// Delete a user
	deleteReq := &pb.DeleteUserRequest{Id: createdUsers[1].GetId()}
	deleteRes, err := c.DeleteUser(ctx, deleteReq)
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}
	log.Printf("User deleted: %v", deleteRes)

	// Get users after deletion to verify
	for _, createdUser := range createdUsers {
		// Make gRPC GetUser request to check if user exists (after deletion)
		getRes, err := c.GetUser(ctx, &pb.GetUserRequest{Id: createdUser.GetId()})
		if err != nil {
			log.Printf("could not get user (likely deleted): %v", err)
		} else {
			log.Printf("User fetched: %v", getRes)
		}
	}
}
