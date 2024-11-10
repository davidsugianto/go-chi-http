package http

import (
	"context"
	entityUser "github.com/davidsugianto/go-chi-http/internal/entity/user"
)

type Service struct {
	Handler
}

type Handler struct {
	userUseCase userUseCase
}

type userUseCase interface {
	CreateUser(ctx context.Context) (data entityUser.User, err error)
}

func New(user userUseCase) *Service {
	return &Service{
		Handler: Handler{
			userUseCase: user,
		},
	}
}
