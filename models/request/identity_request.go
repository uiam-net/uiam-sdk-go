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

// IdentityUpsertRequest IdentityUpsertRequest
type IdentityUpsertRequest struct {
	BaseRequest

	IdentityUUID string                      `json:"uuid"`
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

// IdentityUpdateBasicRequest UpdateRequest
type IdentityUpdateBasicRequest struct {
	IdentityUUID string               `json:"user_id"`
	Username     string               `json:"username"`
	Description  string               `json:"description"`
	AvatarURL    string               `json:"avatar_url"`
	Attributes   uiammodels.Attribute `json:"attributes"`
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
