package server

type Handler struct {
	serverUseCase serverUseCase
}

func New(serverUseCase serverUseCase) *Handler {
	return &Handler{serverUseCase: serverUseCase}
}
