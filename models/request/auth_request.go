package uiamsdk

import (
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// =================== Auth =================== //

// AuthIDCreateRequest AuthIDCreateRequest
type AuthIDCreateRequest struct {
	BaseRequest

	Provider     uiammodel.AuthProviderTypeEnum `json:"provider"`
	IdentityName string                         `json:"username"`
	PhoneCode    string                         `json:"phone_code"`
	PhoneNumber  string                         `json:"phone_number"`
	Email        string                         `json:"email"`
	AvatarURL    string                         `json:"avatar_url"`
	Description  string                         `json:"description"`
	OauthID      string                         `json:"oauth_id" binding:"required,min=4"`
	UnionID      string                         `json:"union_id"`
	AppUserID    string                         `json:"app_user_id"`
	AppUserName  string                         `json:"app_user_name"`
	Credential   uiammodel.Attribute            `json:"credential"`
}

// AuthCreateRequest AuthCreateRequest
type AuthCreateRequest struct {
	BaseRequest

	Provider    uiammodel.AuthProviderTypeEnum `json:"provider"`
	OauthID     string                         `json:"oauth_id" binding:"required,min=4"`
	UnionID     string                         `json:"union_id"`
	AppUserID   string                         `json:"app_user_id"`
	AppUserName string                         `json:"app_user_name"`
	Credential  uiammodel.Attribute            `json:"credential"`
}

// AuthUpdateRequest AuthUpdateRequest
type AuthUpdateRequest struct {
	BaseRequest

	Provider    uiammodel.AuthProviderTypeEnum `json:"provider"`
	OauthID     string                         `json:"oauth_id"`
	UnionID     string                         `json:"union_id"`
	AppUserID   string                         `json:"app_user_id"`
	AppUserName string                         `json:"app_user_name"`
	Credential  uiammodel.Attribute            `json:"credential"`
}

// ============= AuthRequest ============= //

// AuthBindingRequest  auth 绑定
type AuthBindingRequest struct {
	BaseRequest

	UserID      uint64                         `json:"user_id"`
	Provider    uiammodel.AuthProviderTypeEnum `json:"provider"`
	OauthID     string                         `json:"oauth_id"`
	UnionID     string                         `json:"union_id"`
	AppUserName string                         `json:"app_user_name"`
	AppUserID   string                         `json:"app_user_id"`
	Credential  uiammodel.Attribute            `json:"credential"`
}

// AuthVerifyRequest AuthVerifyRequest
type AuthVerifyRequest struct {
	BaseRequest

	Scheme      uiammodel.AuthSchemeEnum `json:"scheme"` // 如果没有 Scheme，就按 Raw 解析
	Credentials string                   `json:"credentials"`
	RawAuth     string                   `json:"raw_auth"`
}

// AuthListRequest 用户表求
type AuthListRequest struct {
	BaseRequest

	Offset   int                            `form:"offset" binding:"gte=0"`
	Limit    int                            `form:"limit" binding:"required,gte=1,lte=500"`
	Provider uiammodel.AuthProviderTypeEnum `form:"provider"`
}
