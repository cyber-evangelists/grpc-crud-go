package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	pb "swag-grpc-crud/proto"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	db *sql.DB
}

func NewUserServiceServer(db *sql.DB) *UserServiceServer {
	return &UserServiceServer{db: db}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	name := req.GetName()
	email := req.GetEmail()

	// Validation: Check if name and email are not empty
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}

	// Validation: Check if name format is valid (starts with an alphabet, followed by alphanumeric characters)
	nameRegex := `^[a-zA-Z][a-zA-Z0-9]*$`
	if matched, _ := regexp.MatchString(nameRegex, name); !matched {
		return nil, errors.New("invalid name format")
	}

	// Validation: Check if email format is valid
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, email); !matched {
		return nil, errors.New("invalid email format")
	}

	var id int
	err := s.db.QueryRowContext(ctx, "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", name, email).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return &pb.CreateUserResponse{Id: fmt.Sprintf("%d", id), Name: name, Email: email}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var id, name, email string
	err := s.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = $1", req.GetId()).Scan(&id, &name, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}
	return &pb.GetUserResponse{Id: id, Name: name, Email: email}, nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	name := req.GetName()
	email := req.GetEmail()

	// Validation: Check if name and email are not empty
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}

	// Validation: Check if name format is valid (starts with an alphabet, followed by alphanumeric characters)
	nameRegex := `^[a-zA-Z][a-zA-Z0-9]*$`
	if matched, _ := regexp.MatchString(nameRegex, name); !matched {
		return nil, errors.New("invalid name format")
	}

	// Validation: Check if email format is valid
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, email); !matched {
		return nil, errors.New("invalid email format")
	}

	// Check if the user exists
	var existingName, existingEmail string
	err := s.db.QueryRowContext(ctx, "SELECT name, email FROM users WHERE id = $1", req.GetId()).Scan(&existingName, &existingEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to check user existence: %v", err)
	}

	// Update user
	_, err = s.db.ExecContext(ctx, "UPDATE users SET name = $1, email = $2 WHERE id = $3", name, email, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return &pb.UpdateUserResponse{Id: req.GetId(), Name: name, Email: email}, nil
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	// Check if the user exists
	var existingID string
	err := s.db.QueryRowContext(ctx, "SELECT id FROM users WHERE id = $1", req.GetId()).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to check user existence: %v", err)
	}

	// Delete user
	_, err = s.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", req.GetId())
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}
	return &pb.DeleteUserResponse{Id: req.GetId()}, nil
}
