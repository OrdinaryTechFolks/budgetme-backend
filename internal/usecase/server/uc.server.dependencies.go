package server

import (
	ServerEntity "github.com/TheUnderdogFolks/budgetme-backend/internal/entity/server"
)

type serverRepo interface {
	GetHealth() (*ServerEntity.Health, error)
}
