package uiamsdk

import (
	"context"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	httputil "github.com/uiam-net/uiam-sdk-go/utils/http"
)

// IdentityRequest IdentityRequest
type IdentityRequest struct {
	Authorization string `json:"token"`
	ServerURL     string `json:"server_url"`
	RealmID       string `json:"realm_id"`
}

// NewUserRequestJwt NewUserRequestJwt
func NewRequestJwt(token, serverURL, realmID string) *IdentityRequest {
	userReq := &IdentityRequest{
		Authorization: fmt.Sprintf("Bearer %s", token),
		ServerURL:     serverURL,
		RealmID:       realmID,
	}

	return userReq
}

// NewUserRequestJwt NewUserRequestJwt
func NewRequest(serverURL string) *IdentityRequest {
	userReq := &IdentityRequest{
		ServerURL: serverURL,
	}

	return userReq
}

// ============ private ============= //
// ============ private ============= //

func (r IdentityRequest) getRequest(ctx context.Context) *resty.Request {
	return httputil.NewRequest(ctx).
		SetHeader("Authorization", r.Authorization).
		SetHeader(httputil.RequestIDKey, httputil.GenRequestID(ctx)).
		SetHeader("UIAM-Realm-ID", r.RealmID)
}
