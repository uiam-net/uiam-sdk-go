package uiamsdk

import (
	"context"
	"fmt"
	"strings"

	resty "github.com/go-resty/resty/v2"
	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
)

// IdentityRequest IdentityRequest
type IdentityRequest struct {
	Authorization string `json:"token"`
	ServerURL     string `json:"server_url"`
}

// NewUserRequestJwt NewUserRequestJwt
func NewUserRequestJwt(token, serverURL string) *IdentityRequest {
	userReq := &IdentityRequest{
		Authorization: fmt.Sprintf("Bearer %s", token),
		ServerURL:     serverURL,
	}

	return userReq
}

// NewUserRequestJwt NewUserRequestJwt
func NewUserRequest(serverURL string) *IdentityRequest {
	userReq := &IdentityRequest{
		ServerURL: serverURL,
	}

	return userReq
}

// ============ api ============= //
// ============ api ============= //

// GetCurrentIdentity GetUser
func (r IdentityRequest) GetCurrentIdentity(ctx context.Context, profile bool) (*uiamresp.Identity, error) {
	var res uiamresp.Identity

	var expand = make([]string, 0)
	if profile {
		expand = append(expand, "profile")
	}

	url := fmt.Sprintf("%s/uv1/identity?expand=%s", r.ServerURL, strings.Join(expand, ","))

	// Request
	if err := Execute(r.getRequest(ctx), "GET", url, nil, &res); nil != err {
		return nil, err
	}

	return &res, nil
}

// ============ private ============= //
// ============ private ============= //

func (r IdentityRequest) getRequest(ctx context.Context) *resty.Request {
	return NewRequest(ctx).
		SetHeader("Authorization", r.Authorization).
		SetHeader(RequestIDKey, GenRequestID(ctx))
}
