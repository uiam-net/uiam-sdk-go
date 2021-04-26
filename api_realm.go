package uiamsdk

import (
	"context"
	"encoding/base64"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
)

// RealmRequest RealmRequest
type RealmRequest struct {
	AuthValue string `json:"auth"`
	ServerURL string `json:"host"`
}

// NewRealmRequestBasic NewRealmRequestBasic
func NewRealmRequestBasic(authKey, authSecret, serverURL string) *RealmRequest {
	ks := authKey + ":" + authSecret
	ksStr := base64.StdEncoding.EncodeToString([]byte(ks))

	appReq := &RealmRequest{
		AuthValue: fmt.Sprintf("Basic %s", ksStr),
		ServerURL: serverURL,
	}

	return appReq
}

// NewRealmRequestJwt NewRealmRequestJwt
func NewRealmRequestJwt(token, serverURL string) *RealmRequest {
	id := &RealmRequest{
		AuthValue: fmt.Sprintf("Bearer %s", token),
		ServerURL: serverURL,
	}

	return id
}

// ============ api ============= //
// ============ api ============= //

// GetCurrentRealm GetCurrentRealm
func (ir RealmRequest) GetCurrentRealm(ctx context.Context) (*uiamresp.Realm, error) {
	var res uiamresp.Realm

	url := fmt.Sprintf("%s/v1/realm", ir.ServerURL)

	// Request
	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &res); nil != err {
		return nil, err
	}

	return &res, nil
}

// ============ private ============= //
// ============ private ============= //

func (ir RealmRequest) getRequest(ctx context.Context) *resty.Request {
	return NewRequest(ctx).
		SetHeader("Authorization", ir.AuthValue).
		SetHeader(RequestIDKey, GenRequestID(ctx))
}
