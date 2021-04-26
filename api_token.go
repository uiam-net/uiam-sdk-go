package uiamsdk

import (
	"context"
	"fmt"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
	uiamreq "github.com/uiam-net/uiam-sdk-go/requests"
	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
)

// ============ api ============= //
// ============ api ============= //

// GenToken GenToken
func (ir IdentityRequest) GenToken(ctx context.Context, req *uiamreq.TokenCreateRequest) (*uiammodel.Token, error) {
	var tokenRes uiammodel.Token
	var url = fmt.Sprintf("%s/v1/identities/%v/tokens", ir.ServerURL, req.Audience)

	if err := Execute(ir.getRequest(ctx), "POST", url, req, &tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}

// GenMfaPhoneCode GenMfaPhoneCode
func (ir IdentityRequest) GenMfaPhoneCode(ctx context.Context, authReq *uiamreq.PhoneCodeVerifyRequest) (string, error) {
	var result map[string]interface{}
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/mfa/phone"), authReq, &result); err != nil {
		return "", err
	}

	if result["code"] == nil {
		return "", uiammodel.NewAppError("result error!")
	}

	return result["code"].(string), nil
}

// VerifyMfaPhoneCode VerifyMfaPhoneCode
func (ir IdentityRequest) VerifyMfaPhoneCode(ctx context.Context, authReq *uiamreq.PhoneCodeVerifyRequest) (*uiamresp.Identity, error) {
	var user uiamresp.Identity
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/mfa/phone/verify"), authReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
