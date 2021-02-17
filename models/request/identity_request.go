package uiamsdk

import (
	uiammodels "github.com/uiam-net/uiam-sdk-go/models"
)

// IdentityListRequest 用户表求
type IdentityListRequest struct {
	BaseRequest

	Offset        int                             `form:"offset" binding:"gte=0"`
	Limit         int                             `form:"limit" binding:"gte=0,lte=500"`
	PhoneCode     string                          `form:"phone_code"`
	PhoneNumber   string                          `form:"phone_number"`
	AppIdentityID string                          `form:"app_user_id"`
	Provider      uiammodels.AuthProviderTypeEnum `form:"provider"`
	Batch         string                          `form:"batch"`
	Expand        string                          `form:"expand"`
}

// IdentityCreateRequest IdentityRequest
type IdentityCreateRequest struct {
	BaseRequest

	IdentityName string                      `json:"username" binding:"required,min=1"`
	Type         uiammodels.IdentityTypeEnum `json:"type" binding:"required,min=1"`
	PhoneCode    string                      `json:"phone_code" binding:"required,min=1"`
	PhoneNumber  string                      `json:"phone_number" binding:"required,min=1"`
	Email        string                      `json:"email"`
	Password     string                      `json:"password"`
	AvatarURL    string                      `json:"avatar_url"`
	Description  string                      `json:"description"`
	Attributes   uiammodels.Attribute        `json:"attributes"`
}

// IdentityChangePhoneRequest IdentityRequest
type IdentityChangePhoneRequest struct {
	BaseRequest

	PhoneCode   string `json:"phone_code" binding:"required,min=1"`
	PhoneNumber string `json:"phone_number" binding:"required,min=1"`
	Email       string `json:"email"`
}

// IdentityUpdateRequest IdentityUpdateRequest
type IdentityUpdateRequest struct {
	BaseRequest

	IdentityName string                      `json:"username" `
	Type         uiammodels.IdentityTypeEnum `json:"type"`
	PhoneCode    string                      `json:"phone_code"`
	PhoneNumber  string                      `json:"phone_number"`
	Email        string                      `json:"email"`
	Password     string                      `json:"password"`
	AvatarURL    string                      `json:"avatar_url"`
	Description  string                      `json:"description"`
	Attributes   uiammodels.Attribute        `json:"attributes"`
}

// =================== Auth =================== //

// AuthMixinCreateRequest AuthMixinCreateRequest
type AuthMixinCreateRequest struct {
	BaseRequest

	IdentityName string               `json:"username"`
	PhoneCode    string               `json:"phone_code" binding:"required,min=1"`
	PhoneNumber  string               `json:"phone_number" binding:"required,min=4"`
	Email        string               `json:"email"`
	AvatarURL    string               `json:"avatar_url"`
	Description  string               `json:"description"`
	MixinID      string               `json:"mixin_id" binding:"required,min=1"`
	MixinName    string               `json:"mixin_name"`
	OauthID      string               `json:"oauth_id" binding:"required,min=4"`
	Credential   uiammodels.Attribute `json:"credential"`
}

// AuthWechatCreateRequest IdentityRequest
type AuthWechatCreateRequest struct {
	BaseRequest

	IdentityName string               `json:"username"`
	PhoneCode    string               `json:"phone_code"`
	PhoneNumber  string               `json:"phone_number"`
	Email        string               `json:"email"`
	AvatarURL    string               `json:"avatar_url"`
	Description  string               `json:"description"`
	OauthID      string               `json:"oauth_id" binding:"required,min=4"`
	UnionID      string               `json:"union_id"`
	Credential   uiammodels.Attribute `json:"credential"`
}