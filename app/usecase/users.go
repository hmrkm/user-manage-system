package usecase

import (
	"context"
	"encoding/json"

	"github.com/hmrkm/user-manage-system/domain"
)

type Users interface {
	List(ctx context.Context, token string, page int, limit int) ([]byte, error)
}

type users struct {
	com            domain.Communicator
	verifyEndpoint string
	baseUrl        string
}

func NewUsers(com domain.Communicator, verifyEndpoint string, baseUrl string) Users {
	return &users{
		com:            com,
		verifyEndpoint: verifyEndpoint,
		baseUrl:        baseUrl,
	}
}

type requestVerify struct {
	Token string
}

func NewRequestVerify(token string) requestVerify {
	return requestVerify{
		Token: token,
	}
}

type responseVerify struct {
	User domain.User
}

type requestUsersList struct {
	Page  int
	Limit int
}

func NewRequestUsersList(page int, limit int) requestUsersList {
	return requestUsersList{
		Page:  page,
		Limit: limit,
	}
}

func (uu *users) List(ctx context.Context, token string, page int, limit int) ([]byte, error) {
	r, err := uu.com.Request(ctx, uu.verifyEndpoint, NewRequestVerify(token))
	if err != nil {
		return nil, domain.ErrNotAuthorized
	}

	res := &responseVerify{}
	if err := json.Unmarshal(r, res); err != nil {
		return nil, domain.ErrJSONUnmarshal
	}

	url := uu.baseUrl + "v1/users/list"
	return uu.com.Request(ctx, url, NewRequestUsersList(page, limit))
}
