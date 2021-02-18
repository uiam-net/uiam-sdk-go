package uiamsdk

import (
	"context"
	"fmt"

	uiammodels "github.com/uiam-net/uiam-sdk-go/models"
	uiamreq "github.com/uiam-net/uiam-sdk-go/models/request"
	uiamresp "github.com/uiam-net/uiam-sdk-go/models/response"
)

// ============ api ============= //
// ============ api ============= //

// // GetAuths GetAuths
// func (ir IdentityRequest) GetAuths(ctx context.Context, provider string, offset, limit int) (*uiammodels.AuthorizationList, error) {
// 	var auths uiammodels.AuthorizationList
// 	var url = fmt.Sprintf("%s/v1/auths?provider=%s&limit=%v&offset=%v", ir.ServerURL, provider, limit, offset)

// 	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auths); err != nil {
// 		return nil, err
// 	}

// 	return &auths, nil
// }

// ============ GET by Identity ============= //

// GetIdentityAuths GetIdentityAuths
func (ir IdentityRequest) GetIdentityAuths(ctx context.Context, uuid string) (*uiamresp.AuthorizationResponse, error) {
	var auth uiamresp.AuthorizationResponse
	var url = fmt.Sprintf("%s/v1/identities/%s/auths", ir.ServerURL, uuid)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// GetIdentityAuth GetIdentityAuth
func (ir IdentityRequest) GetIdentityAuth(ctx context.Context, uuid string, provider uiammodels.AuthProviderTypeEnum) (*uiamresp.AuthorizationResponse, error) {
	var auth uiamresp.AuthorizationResponse
	var url = fmt.Sprintf("%s/v1/identities/%s/auths/%s", ir.ServerURL, uuid, provider)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// ============ GET by OAuthID ============= //

// GetAuthByOAuthID GetAuthByOAuthID
func (ir IdentityRequest) GetAuthByOAuthID(ctx context.Context, provider uiammodels.AuthProviderTypeEnum, oauthID string) (*uiamresp.AuthorizationResponse, error) {
	var auth uiamresp.AuthorizationResponse
	var url = fmt.Sprintf("%s/v1/auths/%s/auths/%s", ir.ServerURL, provider, oauthID)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// ============ REG & Update ============= //

// RegAuthAndID RegAuthAndID
func (ir IdentityRequest) RegAuthAndID(ctx context.Context, req *uiamreq.AuthIDCreateRequest) (*uiamresp.IdentityAuthsResponse, error) {
	var auth uiamresp.IdentityAuthsResponse
	var url = fmt.Sprintf("%s/v1/auths/%s", ir.ServerURL, req.Provider)

	if err := Execute(ir.getRequest(ctx), "POST", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UpdateAuth UpdateAuth
func (ir IdentityRequest) UpdateAuth(ctx context.Context, provider uiammodels.AuthProviderTypeEnum, req *uiamreq.AuthIDCreateRequest) (*uiamresp.AuthorizationResponse, error) {
	var auth uiamresp.AuthorizationResponse
	var url = fmt.Sprintf("%s/v1/auths/%s/auths/%s", ir.ServerURL, provider, req.OauthID)

	if err := Execute(ir.getRequest(ctx), "PATCH", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// ============ ID - AUTH Bind ============= //

// BindAuth BindAuth
func (ir IdentityRequest) BindAuth(ctx context.Context, req *uiamreq.AuthBindingRequest) (*uiammodels.Authorization, error) {
	var auth uiammodels.Authorization
	var url = fmt.Sprintf("%s/v1/identities/%v/auths/%s/bind", ir.ServerURL, req.UserID, req.Provider)

	if err := Execute(ir.getRequest(ctx), "PUT", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UnbindAuth BindAuth
func (ir IdentityRequest) UnbindAuth(ctx context.Context, userID uint64, provider uiammodels.AuthProviderTypeEnum) error {
	var result map[string]interface{}
	var url = fmt.Sprintf("%s/v1/identities/%v/auths/%s/bind", ir.ServerURL, userID, provider)

	if err := Execute(ir.getRequest(ctx), "DELETE", url, nil, &result); err != nil {
		return err
	}

	if result["result"] == nil || result["result"].(string) != "ok" {
		return uiammodels.NewAppError("result error!")
	}

	return nil
}
