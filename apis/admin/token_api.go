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

// GenToken GenToken
func (ir IdentityRequest) GenToken(ctx context.Context, req *uiamreq.TokenCreateRequest) (*uiamresp.LoginResp, error) {
	var tokenRes uiamresp.LoginResp
	var url = fmt.Sprintf("%s/mv1/identities/%v/tokens", ir.ServerURL, req.Audience)

	if err := httputil.Execute(ir.getRequest(ctx), "POST", url, req, &tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}

// GenMfaPhoneCode GenMfaPhoneCode
func (ir IdentityRequest) GenMfaPhoneCode(ctx context.Context, authReq *uiamreq.PhoneCodeVerifyRequest) (string, error) {
	var result map[string]interface{}
	if err := httputil.Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/mv1/mfa/phone"), authReq, &result); err != nil {
		return "", err
	}

	if result["code"] == nil {
		return "", errors.New("result error!")
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
