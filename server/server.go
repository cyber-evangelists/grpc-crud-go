package server

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
	pb "swag-grpc-crud/proto"
)

func StartGRPCServer() {
	db, err := sql.Open("pgx", "postgres://postgres:password@localhost:5432/userdb")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, NewUserServiceServer(db))

	log.Println("gRPC server listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
