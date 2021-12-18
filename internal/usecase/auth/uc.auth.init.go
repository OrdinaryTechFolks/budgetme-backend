package auth

type Auth struct {
	authRepo authRepo
}

func New(authRepo authRepo) *Auth {
	return &Auth{authRepo: authRepo}
}
