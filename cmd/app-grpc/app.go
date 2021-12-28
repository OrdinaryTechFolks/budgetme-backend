package main

import (
	"fmt"

	serverRepo "github.com/TheUnderdogFolks/budgetme-backend/internal/repo/server"
	serverUseCase "github.com/TheUnderdogFolks/budgetme-backend/internal/usecase/server"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func startApp() {
	metrics, err := newrelic.NewApplication(
		newrelic.ConfigAppName("budgetme-backend"),
		newrelic.ConfigLicense("5d45054992c7f69d5dd9bda28b4005f676a4NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		panic(fmt.Sprintf("Failed to initialize newrelic app. %v", err))
	}

	serverRepo := serverRepo.New()
	serverUseCase := serverUseCase.New(serverRepo)

	startServer(metrics, serverUseCase)
}
