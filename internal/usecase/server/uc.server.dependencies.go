package server

import (
	ServerEntity "github.com/OrdinaryTechFolks/budgetme-backend/internal/entity/server"
)

type serverRepo interface {
	GetHealth() (*ServerEntity.Health, error)
}
