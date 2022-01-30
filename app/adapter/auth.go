package adapter

import (
	"context"
	"encoding/json"

	"github.com/hmrkm/user-manage-system/domain"
	"github.com/hmrkm/user-manage-system/usecase"
)

type Auth interface {
	Auth(ctx context.Context, req *RequestAuth) (*ResponseAuth, error)
}

type auth struct {
	usecase usecase.Auth
}

func NewAuth(usecase usecase.Auth) Auth {
	return &auth{
		usecase: usecase,
	}
}

func (aa *auth) Auth(ctx context.Context, req *RequestAuth) (*ResponseAuth, error) {
	r, err := aa.usecase.Auth(ctx, req.Email, req.Password)

	res := &ResponseAuth{}
	if err := json.Unmarshal(r, res); err != nil {
		return nil, domain.ErrJSONUnmarshal
	}

	return res, err
}
