package server

import (
	ServerEntity "github.com/TheWokeDeveloper/budgetme-backend/internal/entity/server"
)

func (s *Server) GetHealth() (*ServerEntity.Health, error) {
	return s.serverRepo.GetHealth()
}
