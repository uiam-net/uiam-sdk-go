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

// IdentityCreateRequest IdentityCreateRequest
type IdentityCreateRequest struct {
	BaseRequest

	Username    string                     `json:"username"`
	Nickname    string                     `json:"nickname"`
	Type        uiammodel.IdentityTypeEnum `json:"_"`
	BizType     string                     `json:"biz_type"`
	PhoneCode   string                     `json:"phone_code"`
	PhoneNumber string                     `json:"phone_number"`
	Email       string                     `json:"email"`
	Password    string                     `json:"password"`
	AvatarURL   string                     `json:"avatar_url"`
	Description string                     `json:"description"`
	Remark      string                     `json:"remark"`
	Attributes  uiammodel.Attribute        `json:"attributes"`
}

// IdentityUpdateRequest IdentityUpdateRequest
type IdentityUpdateRequest struct {
	BaseRequest

	Nickname    string                       `json:"nickname"`
	PhoneCode   string                       `json:"phone_code"`
	PhoneNumber string                       `json:"phone_number"`
	Email       string                       `json:"email"`
	Password    string                       `json:"password"`
	AvatarURL   string                       `json:"avatar_url"`
	Description string                       `json:"description"`
	Remark      string                       `json:"remark"`
	Attributes  uiammodel.Attribute          `json:"attributes"`
	Status      uiammodel.IdentityStatusEnum `json:"status"`
}

// IdentityUpdateBasicRequest UpdateRequest
type IdentityUpdateBasicRequest struct {
	BaseRequest

	Nickname    string              `json:"nickname"`
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
	BaseRequest

	Password string `json:"password"`
}
