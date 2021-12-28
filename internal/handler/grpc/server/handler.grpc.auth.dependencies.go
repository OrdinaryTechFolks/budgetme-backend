package server

import (
	ServerEntity "github.com/TheWokeDeveloper/budgetme-backend/internal/entity/server"
)

type serverUseCase interface {
	GetHealth() (*ServerEntity.Health, error)
}
