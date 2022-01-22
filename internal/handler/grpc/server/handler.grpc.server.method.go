package server

import (
	"context"

	pb "github.com/OrdinaryTechFolks/grpc/budgetme/proto"
)

func (h *Handler) GetHealth(ctx context.Context, req *pb.GetHealthRequest) (*pb.GetHealthResponse, error) {
	var (
		response = &pb.GetHealthResponse{}
	)

	health, err := h.serverUseCase.GetHealth()
	if err != nil {
		return response, err
	}

	return &pb.GetHealthResponse{Message: health.Message}, nil
}
