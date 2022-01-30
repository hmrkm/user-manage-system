package usecase

import (
	"context"

	"github.com/hmrkm/user-manage-system/domain"
)

type Auth interface {
	Auth(ctx context.Context, email string, password string) ([]byte, error)
}

type auth struct {
	com          domain.Communicator
	authEndpoint string
}

func NewAuth(com domain.Communicator, authEndpoint string) Auth {
	return &auth{
		com:          com,
		authEndpoint: authEndpoint,
	}
}

type requestAuth struct {
	Email    string
	Password string
}

func NewRequestAuth(email string, password string) requestAuth {
	return requestAuth{
		Email:    email,
		Password: password,
	}
}

func (au *auth) Auth(ctx context.Context, email string, password string) ([]byte, error) {
	res, err := au.com.Request(ctx, au.authEndpoint, NewRequestAuth(email, password))
	if err != nil {
		return nil, domain.ErrNotAuthorized
	}
	return res, nil
}
