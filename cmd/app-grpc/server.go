package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/TheWokeDeveloper/budgetme-backend/grpc/budgetme/proto"
	serverHandler "github.com/TheWokeDeveloper/budgetme-backend/internal/handler/grpc/server"
	"github.com/TheWokeDeveloper/budgetme-backend/internal/pkg/env"
	serverUseCase "github.com/TheWokeDeveloper/budgetme-backend/internal/usecase/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startServer(serverUseCase *serverUseCase.Server) {
	server := grpc.NewServer()
	reflection.Register(server)

	serverHandler := serverHandler.New(serverUseCase)
	pb.RegisterServerServer(server, serverHandler)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", env.GetInt("port", "8001")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
