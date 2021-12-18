package main

import (
	AuthHandler "github.com/TheUnderdogFolks/budgetme-backend/internal/handler/grpc/auth"
	AuthRepo "github.com/TheUnderdogFolks/budgetme-backend/internal/repo/auth"
	AuthUseCase "github.com/TheUnderdogFolks/budgetme-backend/internal/usecase/auth"
)

func startServer() {
	authRepo := AuthRepo.New()
	authUseCase := AuthUseCase.New(authRepo)
	_ = AuthHandler.New(authUseCase)
}
