package uiamsdk

import (
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// IdentityListRequest 用户表求
type IdentityListRequest struct {
	BaseRequest

	BizType     string                        `form:"type"`
	Offset      int                           `form:"offset" binding:"gte=0"`
	Limit       int                           `form:"limit" binding:"gte=0,lte=500"`
	PhoneCode   string                        `form:"phone_code"`
	PhoneNumber string                        `form:"phone_number"`
	OAuthAppID  string                        `form:"app_user_id"`
	Provider    uiammodel.ConnectProviderEnum `form:"provider"`
	Batch       string                        `form:"batch"`
	Expand      string                        `form:"expand"`
}

// IdentityUpsertRequest IdentityUpsertRequest
type IdentityUpsertRequest struct {
	BaseRequest

	IdentityName string                     `json:"username" binding:"required,min=1"`
	Type         uiammodel.IdentityTypeEnum `json:"_" binding:"required,min=1"`
	BizType      string                     `json:"type" binding:"required,min=1"`
	PhoneCode    string                     `json:"phone_code" binding:"required,min=1"`
	PhoneNumber  string                     `json:"phone_number" binding:"required,min=1"`
	Email        string                     `json:"email"`
	Password     string                     `json:"password"`
	AvatarURL    string                     `json:"avatar_url"`
	Description  string                     `json:"description"`
	Remark       string                     `json:"remark"`
	Attributes   uiammodel.Attribute        `json:"attributes"`
}

// IdentityUpdateRequest IdentityUpdateRequest
type IdentityUpdateRequest struct {
	BaseRequest

	IdentityName string                       `json:"username"`
	PhoneCode    string                       `json:"phone_code"`
	PhoneNumber  string                       `json:"phone_number"`
	Email        string                       `json:"email"`
	Password     string                       `json:"password"`
	AvatarURL    string                       `json:"avatar_url"`
	Description  string                       `json:"description"`
	Remark       string                       `json:"remark"`
	Attributes   uiammodel.Attribute          `json:"attributes"`
	Status       uiammodel.IdentityStatusEnum `json:"status"`
}

// IdentityUpdateBasicRequest UpdateRequest
type IdentityUpdateBasicRequest struct {
	BaseRequest

	PhoneCode   string              `json:"username"`
	Description string              `json:"description"`
	AvatarURL   string              `json:"avatar_url"`
	Attributes  uiammodel.Attribute `json:"attributes"`
}

// IdentityUpdateBasicRequest UpdateRequest
type IdentityMfaRequest struct {
	BaseRequest

	MfaType  uiammodel.MfaTypeEnum `json:"type"`
	Passcode string                `json:"passcode"`
}

// ================ Update ================= //

// IdentityChangePhoneRequest IdentityRequest
type IdentityChangePhoneRequest struct {
	BaseRequest

	PhoneCode   string `json:"phone_code" binding:"required,min=1"`
	PhoneNumber string `json:"phone_number" binding:"required,min=1"`
	Email       string `json:"email"`
}

// UserResetPasswordRequest UpdateRequest
type UserResetPasswordRequest struct {
	UserID   string `json:"user_id"`
	RealmID  string `json:"realm_id"`
	Password string `json:"password"`
}
