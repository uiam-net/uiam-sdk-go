package uiamsdk

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// TokenCreateRequest TokenCreateRequest
type TokenCreateRequest struct {
	BaseRequest

	JTI       string        `json:"jti,omitempty"`
	Audience  string        `json:"aud,omitempty"`
	Issuer    string        `json:"iss,omitempty"`
	NotBefore int64         `json:"nbf,omitempty"`
	Subject   string        `json:"sub,omitempty"`
	Duration  time.Duration `json:"duration"` // ExpriedAt 有效时长

	// Custom
	RealmID      uint64                     `json:"realm_id,omitempty"`
	RealmSlug    string                     `json:"realm_slug,omitempty"`
	MerchantID   uint64                     `json:"merchant_id,omitempty"`
	MerchantSlug string                     `json:"merchant_slug,omitempty"`
	IdentityType uiammodel.IdentityTypeEnum `json:"identity_type,omitempty"`
	Scheme       uiammodel.AuthSchemeEnum   `json:"scheme,omitempty"`      // 签名方式
	UID          string                     `json:"uid,omitempty"`         // Audience 区别
	Device       string                     `json:"device,omitempty"`      // 每个 device 会有一个相对固定 session
	Description  string                     `json:"description,omitempty"` // device 介绍
	SessionID    string                     `json:"sid,omitempty"`         // 使用哪个 session 进行签名
	Sign         string                     `json:"sig,omitempty"`         // ?
	SignAlg      string                     `json:"sal,omitempty"`         // ?
	Extra        string                     `json:"extra,omitempty"`       // 签名时，会把这部分也签进去
}
