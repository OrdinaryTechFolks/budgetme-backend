package server

import "github.com/OrdinaryTechFolks/budgetme-backend/internal/config"

type Repo struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Repo {
	return &Repo{cfg: cfg}
}
