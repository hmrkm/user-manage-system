package adapter

import (
	"context"
	"encoding/json"

	"github.com/hmrkm/user-manage-system/domain"
	"github.com/hmrkm/user-manage-system/usecase"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/user-manage-system/$GOPACKAGE -package=$GOPACKAGE -destination=users_mock.go
type Users interface {
	List(ctx context.Context, req *RequestUsersList) (*ResponseUsersList, error)
}

type users struct {
	usecase usecase.Users
}

func NewUsers(usecase usecase.Users) Users {
	return &users{
		usecase: usecase,
	}
}

func (ua *users) List(ctx context.Context, req *RequestUsersList) (*ResponseUsersList, error) {
	r, err := ua.usecase.List(ctx, req.AuthToken, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	res := &ResponseUsersList{}
	if err := json.Unmarshal(r, res); err != nil {
		return nil, domain.ErrJSONUnmarshal
	}

	return res, nil
}
