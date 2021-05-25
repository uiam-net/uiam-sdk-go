package uiamsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
	uiamreq "github.com/uiam-net/uiam-sdk-go/requests"
	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
	httputil "github.com/uiam-net/uiam-sdk-go/utils/http"
)

// ============ api ============= //
// ============ api ============= //

// GenToken GenToken
func (ir IdentityRequest) Login(ctx context.Context, req *uiamreq.LoginRequest) (*uiamresp.LoginResp, error) {
	tokenRes := new(uiamresp.LoginResp)

	var result map[string]interface{}
	var url = fmt.Sprintf("%s/mv1/login", ir.ServerURL)

	if err := httputil.Execute(ir.getRequest(ctx), "POST", url, req, &result); err != nil {
		return nil, err
	}

	// fmt.Printf("======result: %v \n", result)
	loginType := result["type"]

	if loginType == nil {
		return nil, errors.New("result error")
	}

	// fmt.Printf("======loginType: %v \n", loginType)
	if loginType == uiammodel.SessionTypeEnumToken.String() {
		loginToken := new(uiamresp.LoginToken)
		payloadJson, _ := json.Marshal(result)
		if err := json.Unmarshal([]byte(payloadJson), loginToken); err != nil {
			return nil, errors.New("unmarshal error")
		}

		tokenRes.Type = uiammodel.SessionTypeEnumToken
		tokenRes.Token = loginToken

		return tokenRes, nil
	}

	return nil, errors.New("result error")
}
