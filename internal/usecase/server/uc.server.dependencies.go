package server

import (
	ServerEntity "github.com/TheWokeDeveloper/budgetme-backend/internal/entity/server"
)

type serverRepo interface {
	GetHealth() (*ServerEntity.Health, error)
}
