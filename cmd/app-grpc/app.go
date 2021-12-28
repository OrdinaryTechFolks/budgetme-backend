package main

import (
	serverRepo "github.com/TheWokeDeveloper/budgetme-backend/internal/repo/server"
	serverUseCase "github.com/TheWokeDeveloper/budgetme-backend/internal/usecase/server"
)

func startApp() {
	serverRepo := serverRepo.New()

	serverUseCase := serverUseCase.New(serverRepo)

	startServer(serverUseCase)
}
