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
	rightsEndpoint string
	baseURL        string
	apiPath        string
}

func (u *users) getApiURL() string {
	return u.baseURL + u.apiPath
}

func NewUsers(com domain.Communicator, verifyEndpoint string, rightsEndpoint string, baseURL string, apiPath string) Users {
	return &users{
		com:            com,
		verifyEndpoint: verifyEndpoint,
		rightsEndpoint: rightsEndpoint,
		baseURL:        baseURL,
		apiPath:        apiPath,
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

type responseAuth struct {
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

type requestRights struct {
	UserId   string
	Resource string
}

func NewRequestRights(userId string, resource string) requestRights {
	return requestRights{
		UserId:   userId,
		Resource: resource,
	}
}

func (uu *users) List(ctx context.Context, token string, page int, limit int) ([]byte, error) {
	authResp, err := uu.com.Request(ctx, uu.verifyEndpoint, NewRequestVerify(token))
	if err != nil {
		return nil, domain.ErrNotAuthorized
	}

	authResponse := &responseAuth{}
	if err := json.Unmarshal(authResp, authResponse); err != nil {
		return nil, domain.ErrJSONUnmarshal
	}

	if _, err := uu.com.Request(ctx, uu.rightsEndpoint, NewRequestRights(authResponse.User.ID, uu.apiPath)); err != nil {
		return nil, domain.ErrForbidden
	}

	return uu.com.Request(ctx, uu.getApiURL(), NewRequestUsersList(page, limit))
}
