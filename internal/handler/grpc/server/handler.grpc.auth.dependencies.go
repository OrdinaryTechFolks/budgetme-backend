package server

import (
	ServerEntity "github.com/OrdinaryTechFolks/budgetme-backend/internal/entity/server"
)

type serverUseCase interface {
	GetHealth() (*ServerEntity.Health, error)
}
