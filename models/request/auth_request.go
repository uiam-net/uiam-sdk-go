package uiamsdk

import (
	uiammodels "github.com/uiam-net/uiam-sdk-go/models"
)

// =================== Auth =================== //

// AuthIDCreateRequest AuthIDCreateRequest
type AuthIDCreateRequest struct {
	BaseRequest

	Provider     uiammodels.AuthProviderTypeEnum `json:"provider"`
	IdentityName string                          `json:"username"`
	PhoneCode    string                          `json:"phone_code"`
	PhoneNumber  string                          `json:"phone_number"`
	Email        string                          `json:"email"`
	AvatarURL    string                          `json:"avatar_url"`
	Description  string                          `json:"description"`
	OauthID      string                          `json:"oauth_id" binding:"required,min=4"`
	UnionID      string                          `json:"union_id"`
	AppUserID    string                          `json:"app_user_id"`
	AppUserName  string                          `json:"app_user_name"`
	Credential   uiammodels.Attribute            `json:"credential"`
}

// AuthCreateRequest AuthCreateRequest
type AuthCreateRequest struct {
	BaseRequest

	Provider    uiammodels.AuthProviderTypeEnum `json:"provider"`
	OauthID     string                          `json:"oauth_id" binding:"required,min=4"`
	UnionID     string                          `json:"union_id"`
	AppUserID   string                          `json:"app_user_id"`
	AppUserName string                          `json:"app_user_name"`
	Credential  uiammodels.Attribute            `json:"credential"`
}
