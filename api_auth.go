package uiamsdk

import (
	"context"
	"fmt"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
	uiamreq "github.com/uiam-net/uiam-sdk-go/models/request"
	uiamresp "github.com/uiam-net/uiam-sdk-go/models/response"
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
func (ir IdentityRequest) GetIdentityAuths(ctx context.Context, uuid string) (*uiamresp.AuthorizationResponse, error) {
	var auth uiamresp.AuthorizationResponse
	var url = fmt.Sprintf("%s/v1/identities/%s/auths", ir.ServerURL, uuid)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// GetIdentityAuth GetIdentityAuth
func (ir IdentityRequest) GetIdentityAuth(ctx context.Context, uuid string, provider uiammodel.ConnectProviderTypeEnum) (*uiamresp.AuthorizationResponse, error) {
	var auth uiamresp.AuthorizationResponse
	var url = fmt.Sprintf("%s/v1/identities/%s/auths/%s", ir.ServerURL, uuid, provider)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// ============ GET by OAuthID ============= //

// GetAuthByOAuthID GetAuthByOAuthID
func (ir IdentityRequest) GetAuthByOAuthID(ctx context.Context, provider uiammodel.ConnectProviderTypeEnum, oauthID string) (*uiamresp.AuthorizationResponse, error) {
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

// UpdateAuthByID UpdateAuthByID
func (ir IdentityRequest) UpdateAuthByID(ctx context.Context, req *uiamreq.AuthUpdateRequest) (*uiamresp.AuthorizationResponse, error) {
	var auth uiamresp.AuthorizationResponse

	var url = fmt.Sprintf("%s/v1/identities/%s/auths/%s", ir.ServerURL, req.IdentityID, req.Provider)

	// fmt.Printf("#INFO = IDAuthorizationupdate: %v", req.AppUserName)
	// fmt.Printf("#INFO = IDAuthorizationupdate: %v", req.UnionID)

	if err := Execute(ir.getRequest(ctx), "PATCH", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UpdateAuthByOauthID UpdateAuth
func (ir IdentityRequest) UpdateAuthByOauthID(ctx context.Context, req *uiamreq.AuthIDCreateRequest) (*uiamresp.AuthorizationResponse, error) {
	var auth uiamresp.AuthorizationResponse

	var url = fmt.Sprintf("%s/v1/auths/%s/auths/%s", ir.ServerURL, req.Provider, req.OauthID)

	if err := Execute(ir.getRequest(ctx), "PATCH", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// ============ ID - AUTH Bind ============= //

// BindAuth BindAuth
func (ir IdentityRequest) BindAuth(ctx context.Context, req *uiamreq.AuthBindingRequest) (*uiammodel.Connect, error) {
	var auth uiammodel.Connect
	var url = fmt.Sprintf("%s/v1/identities/%v/auths/%s/bind", ir.ServerURL, req.UserID, req.Provider)

	if err := Execute(ir.getRequest(ctx), "PUT", url, req, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}

// UnbindAuth BindAuth
func (ir IdentityRequest) UnbindAuth(ctx context.Context, userID uint64, provider uiammodel.ConnectProviderTypeEnum) error {
	var result map[string]interface{}
	var url = fmt.Sprintf("%s/v1/identities/%v/auths/%s/bind", ir.ServerURL, userID, provider)

	if err := Execute(ir.getRequest(ctx), "DELETE", url, nil, &result); err != nil {
		return err
	}

	if result["result"] == nil || result["result"].(string) != "ok" {
		return uiammodel.NewAppError("result error!")
	}

	return nil
}
