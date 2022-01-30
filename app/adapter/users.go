package adapter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hmrkm/user-manage-system/domain"
)

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/user-manage-system/$GOPACKAGE -package=$GOPACKAGE -destination=users_mock.go
type UsersAdapter interface {
	List(ctx context.Context, req *RequestUsersList, token string) (*ResponseUsersList, error)
}

type usersAdapter struct {
	com          domain.Communicator
	authEndpoint string
}

func NewUserAdapter(com domain.Communicator, authEndpoint string) UsersAdapter {
	return &usersAdapter{
		com:          com,
		authEndpoint: authEndpoint,
	}
}

func (ua *usersAdapter) List(ctx context.Context, req *RequestUsersList, token string) (*ResponseUsersList, error) {
	res, err := ua.com.Request(ctx, ua.authEndpoint, map[string]string{
		"token": token,
	})
	if err != nil {
		return nil, domain.ErrNotAuthorized
	}
	user := &User{}
	json.Unmarshal(res, user)
	fmt.Println(user)

	return &ResponseUsersList{}, nil
}
