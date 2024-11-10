package user

import (
	"context"
	entityUser "github.com/davidsugianto/go-chi-http/internal/entity/user"
)

func (u *UseCase) CreateUser(ctx context.Context) (data entityUser.User, err error) {
	return data, nil
}
