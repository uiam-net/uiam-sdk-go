package uiamsdk

import (
	"context"
	"errors"
	"fmt"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
	uiamreq "github.com/uiam-net/uiam-sdk-go/requests"
	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
	httputil "github.com/uiam-net/uiam-sdk-go/utils/http"
)

// ============ api ============= //
// ============ api ============= //

// // GetAuths GetAuths
// func (ir IdentityRequest) GetAuths(ctx context.Context, provider string, offset, limit int) (*uiammodel.AuthorizationList, error) {
// 	var auths uiammodel.AuthorizationList
// 	var url = fmt.Sprintf("%s/v1/auths?provider=%s&limit=%v&offset=%v", ir.ServerURL, provider, limit, offset)

// 	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auths); err != nil {
// 		return nil, err
// 	}

// 	return &auths, nil
// }

// ============ GET by Identity ============= //

// GetIdentityAuths GetIdentityAuths
func (ir IdentityRequest) GetIdentityAuths(ctx context.Context, uuid string) (*uiamresp.Authorization, error) {
	var auth uiamresp.Authorization
	var url = fmt.Sprintf("%s/mv1/identities/%s/auths", ir.ServerURL, uuid)

	if err := httputil.Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// GetIdentityAuth GetIdentityAuth
func (ir IdentityRequest) GetIdentityAuth(ctx context.Context, uuid string, provider uiammodel.ConnectProviderEnum) (*uiamresp.Authorization, error) {
	var auth uiamresp.Authorization
	var url = fmt.Sprintf("%s/mv1/identities/%s/auths/%s", ir.ServerURL, uuid, provider)

	if err := httputil.Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// ============ GET by OAuthID ============= //

// GetAuthByOAuthID GetAuthByOAuthID
func (ir IdentityRequest) GetAuthByOAuthID(ctx context.Context, provider uiammodel.ConnectProviderEnum, oauthID string) (*uiamresp.Authorization, error) {
	var auth uiamresp.Authorization
	var url = fmt.Sprintf("%s/mv1/auths/%s/auths/%s", ir.ServerURL, provider, oauthID)

	if err := httputil.Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// ============ REG & Update ============= //

// RegAuthAndID RegAuthAndID
func (ir IdentityRequest) RegAuthAndID(ctx context.Context, req *uiamreq.AuthIDCreateRequest) (*uiamresp.Identity, error) {
	var auth uiamresp.Identity
	var url = fmt.Sprintf("%s/mv1/auths/%s", ir.ServerURL, req.Provider)

	if err := httputil.Execute(ir.getRequest(ctx), "POST", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UpdateAuthByID UpdateAuthByID
func (ir IdentityRequest) UpdateAuthByID(ctx context.Context, req *uiamreq.AuthUpdateRequest) (*uiamresp.Authorization, error) {
	var auth uiamresp.Authorization

	var url = fmt.Sprintf("%s/mv1/identities/%s/auths/%s", ir.ServerURL, req.IdentityUUID, req.Provider)

	if err := httputil.Execute(ir.getRequest(ctx), "PATCH", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UpdateAuthByOauthID UpdateAuth
func (ir IdentityRequest) UpdateAuthByOauthID(ctx context.Context, req *uiamreq.AuthIDCreateRequest) (*uiamresp.Authorization, error) {
	var auth uiamresp.Authorization

	var url = fmt.Sprintf("%s/mv1/auths/%s/auths/%s", ir.ServerURL, req.Provider, req.OauthID)

	if err := httputil.Execute(ir.getRequest(ctx), "PATCH", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// ============ ID - AUTH Bind ============= //

// BindAuth BindAuth
func (ir IdentityRequest) BindAuth(ctx context.Context, req *uiamreq.AuthBindingRequest) (*uiamresp.Connect, error) {
	var auth uiamresp.Connect
	var url = fmt.Sprintf("%s/mv1/identities/%v/auths/%s/bind", ir.ServerURL, req.UserID, req.Provider)

	if err := httputil.Execute(ir.getRequest(ctx), "PUT", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UnbindAuth BindAuth
func (ir IdentityRequest) UnbindAuth(ctx context.Context, userID uint64, provider uiammodel.ConnectProviderEnum) error {
	var result map[string]interface{}
	var url = fmt.Sprintf("%s/mv1/identities/%v/auths/%s/bind", ir.ServerURL, userID, provider)

	if err := httputil.Execute(ir.getRequest(ctx), "DELETE", url, nil, &result); err != nil {
		return err
	}

	if result["result"] == nil || result["result"].(string) != "ok" {
		return errors.New("result error")
	}

	return nil
}
