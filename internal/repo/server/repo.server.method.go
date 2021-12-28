package server

import (
	"fmt"

	ServerEntity "github.com/TheWokeDeveloper/budgetme-backend/internal/entity/server"
)

func (r *Repo) GetHealth() (*ServerEntity.Health, error) {
	return &ServerEntity.Health{Message: fmt.Sprintf("Server running on %s", r.cfg.Version)}, nil
}
