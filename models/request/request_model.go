package uiamsdk

import (
	"time"

	uiammodels "github.com/uiam-net/uiam-sdk-go/models"
)

// CreateUserReq CreateUserReq
type CreateUserReq struct {
	Username    string               `json:"username"`
	Type        string               `json:"type"`
	PhoneCode   string               `json:"phone_code"`
	PhoneNumber string               `json:"phone_number"`
	Email       string               `json:"email"`
	Password    string               `json:"password"`
	AvatarURL   string               `json:"avatar_url"`
	Description string               `json:"description"`
	Attributes  uiammodels.Attribute `json:"attributes"`
}

// UpdateBasicUpdateRequest UpdateRequest
type UpdateBasicUpdateRequest struct {
	UserID      string               `json:"user_id"`
	Username    string               `json:"username"`
	Description string               `json:"description"`
	AvatarURL   string               `json:"avatar_url"`
	Attributes  uiammodels.Attribute `json:"attributes"`
}

// UserFullUpdateRequest UserModifyReq
type UserFullUpdateRequest struct {
	UserID      uint64               `json:"user_id"`
	Username    string               `json:"username"`
	Type        string               `json:"type"`
	PhoneCode   string               `json:"phone_code"`
	PhoneNumber string               `json:"phone_number"`
	Email       string               `json:"email"`
	Password    string               `json:"password"`
	AvatarURL   string               `json:"avatar_url"`
	Description string               `json:"description"`
	Attributes  uiammodels.Attribute `json:"attributes"`
}

// UserResetPasswordRequest UpdateRequest
type UserResetPasswordRequest struct {
	UserID   string `json:"user_id"`
	RealmID  string `json:"realm_id"`
	Password string `json:"password"`
}

// KycSyncRequest KycSyncRequest
type KycSyncRequest struct {
	UserID      uint64     `json:"user_id"`
	KycLevel    int        `json:"kyc_level"`
	KycStatus   string     `json:"kyc_status"`
	KycError    string     `json:"kyc_error"`
	CertifiedAt *time.Time `json:"certified_at"`
	IDNo        string     `json:"id_no"`
	Name        string     `json:"name"`
	Country     string     `json:"country"`
	IDType      string     `json:"id_type"`
}

// ============= AuthRequest ============= //

// AuthBindingRequest  auth 绑定
type AuthBindingRequest struct {
	BaseRequest

	UserID      uint64                          `json:"user_id"`
	Provider    uiammodels.AuthProviderTypeEnum `json:"provider"`
	OauthID     string                          `json:"oauth_id"`
	UnionID     string                          `json:"union_id"`
	AppUserName string                          `json:"app_user_name"`
	AppUserID   string                          `json:"app_user_id"`
	Credential  uiammodels.Attribute            `json:"credential"`
}

// AuthVerifyRequest AuthVerifyRequest
type AuthVerifyRequest struct {
	BaseRequest

	Scheme      uiammodels.AuthSchemeEnum `json:"scheme"` // 如果没有 Scheme，就按 Raw 解析
	Credentials string                    `json:"credentials"`
	RawAuth     string                    `json:"raw_auth"`
}

// AuthListRequest 用户表求
type AuthListRequest struct {
	BaseRequest

	Offset   int                             `form:"offset" binding:"gte=0"`
	Limit    int                             `form:"limit" binding:"required,gte=1,lte=500"`
	Provider uiammodels.AuthProviderTypeEnum `form:"provider"`
}
