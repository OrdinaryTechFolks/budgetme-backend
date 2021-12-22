package server

type Server struct {
	serverRepo serverRepo
}

func New(serverRepo serverRepo) *Server {
	return &Server{serverRepo: serverRepo}
}
