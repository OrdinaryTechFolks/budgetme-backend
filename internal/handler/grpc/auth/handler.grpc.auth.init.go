package auth

type Handler struct {
	authUseCase authUseCase
}

func New(authUseCase authUseCase) *Handler {
	return &Handler{authUseCase: authUseCase}
}
