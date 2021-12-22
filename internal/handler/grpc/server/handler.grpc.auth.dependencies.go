package server

import (
	ServerEntity "github.com/TheUnderdogFolks/budgetme-backend/internal/entity/server"
)

type serverUseCase interface {
	GetHealth() (*ServerEntity.Health, error)
}
