package uiamsdk

import (
	"context"
	"fmt"

	uiammodels "github.com/uiam-net/uiam-sdk-go/models"
	uiamreq "github.com/uiam-net/uiam-sdk-go/models/request"
)

// ============ api ============= //
// ============ api ============= //

// GenToken GenToken
func (ir IdentityRequest) GenToken(ctx context.Context, req *uiamreq.TokenCreateRequest) (*uiammodels.Token, error) {
	var tokenRes uiammodels.Token
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
		return "", uiammodels.NewAppError("result error!")
	}

	return result["code"].(string), nil
}

// VerifyMfaPhoneCode VerifyMfaPhoneCode
func (ir IdentityRequest) VerifyMfaPhoneCode(ctx context.Context, authReq *uiamreq.PhoneCodeVerifyRequest) (*uiammodels.Identity, error) {
	var user uiammodels.Identity
	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/mfa/phone/verify"), authReq, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
