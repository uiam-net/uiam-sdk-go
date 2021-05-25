package uiamsdk

import (
	"context"
	"errors"
	"fmt"

	uiamreq "github.com/uiam-net/uiam-sdk-go/requests"
	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
	httputil "github.com/uiam-net/uiam-sdk-go/utils/http"
)

// ============ api ============= //
// ============ api ============= //

// GenToken
func (ir IdentityRequest) GenToken(ctx context.Context, req *uiamreq.TokenCreateRequest) (*uiamresp.LoginToken, error) {
	var tokenRes uiamresp.LoginToken
	var url = fmt.Sprintf("%s/mv1/identities/%v/token", ir.ServerURL, req.Subject)

	if err := httputil.Execute(ir.getRequest(ctx), "POST", url, req, &tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}

// GenTokenByLogin
func (ir IdentityRequest) GenTokenByLogin(ctx context.Context, req *uiamreq.LoginRequest) (*uiamresp.LoginToken, error) {
	var tokenRes uiamresp.LoginToken
	var url = fmt.Sprintf("%s/mv1/token", ir.ServerURL)

	if err := httputil.Execute(ir.getRequest(ctx), "POST", url, req, &tokenRes); err != nil {
		return nil, err
	}

	fmt.Printf("%v", tokenRes)
	return &tokenRes, nil
}

// GenMfaPhoneCode GenMfaPhoneCode
func (ir IdentityRequest) GenMfaPhoneCode(ctx context.Context, authReq *uiamreq.PhoneCodeVerifyRequest) (string, error) {
	var result map[string]interface{}

	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/mv1/mfa/phone"), authReq, &result); err != nil {
		return "", err
	}

	if result["code"] == nil {
		return "", errors.New("result error")
	}

	return result["code"].(string), nil
}

// VerifyMfaPhoneCode VerifyMfaPhoneCode
func (ir IdentityRequest) VerifyMfaPhoneCode(ctx context.Context, authReq *uiamreq.PhoneCodeVerifyRequest) (*uiamresp.Identity, error) {
	var user uiamresp.Identity

	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/mv1/mfa/phone/verify"), authReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
