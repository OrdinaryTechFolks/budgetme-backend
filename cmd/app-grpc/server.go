package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/TheUnderdogFolks/budgetme-backend/grpc/budgetme/proto"
	serverHandler "github.com/TheUnderdogFolks/budgetme-backend/internal/handler/grpc/server"
	"github.com/TheUnderdogFolks/budgetme-backend/internal/pkg/env"
	serverUseCase "github.com/TheUnderdogFolks/budgetme-backend/internal/usecase/server"
	"github.com/newrelic/go-agent/v3/newrelic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
)

func startServer(metrics *newrelic.Application, serverUseCase *serverUseCase.Server) {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(nrgrpc.UnaryServerInterceptor(metrics)),
		grpc.StreamInterceptor(nrgrpc.StreamServerInterceptor(metrics)),
	)
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
