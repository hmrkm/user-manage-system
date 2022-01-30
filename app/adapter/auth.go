package adapter

import (
	"context"
	"encoding/json"

	"github.com/hmrkm/user-manage-system/domain"
)

type AuthAdapter interface {
	Auth(ctx context.Context, req *RequestAuth) (*ResponseAuth, error)
}

type authAdapter struct {
	com          domain.Communicator
	authEndpoint string
}

func NewAuthAtapdater(com domain.Communicator, authEndpoint string) AuthAdapter {
	return &authAdapter{
		com:          com,
		authEndpoint: authEndpoint,
	}
}

func (aa *authAdapter) Auth(ctx context.Context, req *RequestAuth) (*ResponseAuth, error) {
	res, err := aa.com.Request(ctx, aa.authEndpoint, map[string]string{
		"token": token,
	})
	if err != nil {
		return nil, domain.ErrNotAuthorized
	}
	user := &User{}
	json.Unmarshal(res, user)

	return nil, nil
}
