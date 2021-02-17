package uiamsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	uiammodels "github.com/uiam-net/uiam-sdk-go/models"
	uiamreq "github.com/uiam-net/uiam-sdk-go/models/request"
	uiamresp "github.com/uiam-net/uiam-sdk-go/models/response"
)

// ============ api ============= //
// ============ api ============= //

// GetAllUsers GetAllUsers
func (ir IdentityRequest) GetAllUsers(ctx context.Context) (*uiamresp.BasePageResponse, error) {
	var resp uiamresp.BasePageResponse

	if err := Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/identities", ir.ServerURL), nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetUser GetUser
func (ir IdentityRequest) GetUser(ctx context.Context, userID string, profile bool) (*uiammodels.Identity, error) {
	var resp uiammodels.Identity

	var expand = make([]string, 0)
	if profile {
		expand = append(expand, "profile")
	}

	url := fmt.Sprintf("%s/v1/identities/%s/?expand=%s", ir.ServerURL, userID, strings.Join(expand, ","))

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateUser GetUser
func (ir IdentityRequest) UpdateUser(ctx context.Context, req *uiamreq.UpdateBasicUpdateRequest) (*uiammodels.Identity, error) {
	var resp uiammodels.Identity

	url := fmt.Sprintf("%s/v1/identities/%s", ir.ServerURL, req.UserID)

	if err := Execute(ir.getRequest(ctx), "PATCH", url, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUsersByIdentityIDs GetAllUsers
func (ir IdentityRequest) GetUsersByIdentityIDs(ctx context.Context, identityIDs []string) ([]*uiammodels.Identity, error) {
	var users []*uiammodels.Identity

	if err := Execute(ir.getRequest(ctx), "GET", fmt.Sprintf("%s/v1/identities?id=%s", ir.ServerURL, strings.Join(identityIDs, ",")), nil, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByPhone GetUser
func (ir IdentityRequest) GetUserByPhone(ctx context.Context, phoneCode, phoneNumber string) (*uiammodels.Identity, error) {
	var resp uiamresp.BasePageResponse

	url := fmt.Sprintf("%s/v1/identities?phone_code=%s&phone_number=%s&limit=1", ir.ServerURL, phoneCode, phoneNumber)

	if err := Execute(ir.getRequest(ctx), "GET", url, nil, &resp); err != nil {
		return nil, err
	}

	if len(resp.Items) > 0 {
		userBt, err := json.Marshal(resp.Items[0])
		if err != nil {
			return nil, uiammodels.NewAppError(err.Error())
		}

		user := new(uiammodels.Identity)
		err2 := json.Unmarshal(userBt, user)
		if err2 != nil {
			return nil, uiammodels.NewAppError(err2.Error())
		}
		return user, nil
	}

	return nil, nil
}

// VerifyUserPassword VerifyUserPassword
func (ir IdentityRequest) VerifyUserPassword(ctx context.Context, identityID string, password string) (*uiammodels.Identity, error) {
	var resp uiammodels.Identity

	url := fmt.Sprintf("%s/v1/identities/%v/password/verify", ir.ServerURL, identityID)

	if err := Execute(ir.getRequest(ctx), "POST", url, map[string]string{"password": password}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateUser CreateUser
func (ir IdentityRequest) CreateUser(ctx context.Context, req *uiamreq.CreateUserReq) (*uiammodels.Identity, error) {
	var user uiammodels.Identity

	if err := Execute(ir.getRequest(ctx), "POST", fmt.Sprintf("%s%s", ir.ServerURL, "/v1/identities"), req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// ChangePassword ChangePassword
func (ir IdentityRequest) ChangePassword(ctx context.Context, req *uiamreq.UserResetPasswordRequest) (*uiammodels.Identity, error) {
	var user uiammodels.Identity

	url := fmt.Sprintf("%s/v1/identities/%v/password", ir.ServerURL, req.UserID)

	if err := Execute(ir.getRequest(ctx), "PUT", url, req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// ChangePhone ChangePhone
func (ir IdentityRequest) ChangePhone(ctx context.Context, req *uiamreq.UserFullUpdateRequest) (*uiammodels.Identity, error) {
	var user uiammodels.Identity

	url := fmt.Sprintf("%s/v1/identities/%d/phone", ir.ServerURL, req.UserID)

	if err := Execute(ir.getRequest(ctx), "PUT", url, req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
