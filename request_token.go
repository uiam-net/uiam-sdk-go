package uiamsdk

import (
	"time"
)

// TokenCreateRequest TokenCreateRequest
type TokenCreateRequest struct {
	BaseRequest

	Scheme AuthSchemeEnum `json:"scheme"`

	Audience  string `json:"aud,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`

	// Custom
	Provider  AuthProviderTypeEnum `json:"oap,omitempty"`
	UID       string               `json:"uid,omitempty"`
	SessionID string               `json:"sid,omitempty"`
	Sign      string               `json:"sig,omitempty"`
	SignAlg   string               `json:"sal,omitempty"`
	Extra     string               `json:"extra,omitempty"`

	// ExpriedAt
	Duration time.Duration `json:"duration"` //  有效时长
}

// // TokenCreateRequest TokenCreateRequest
// type TokenCreateRequest struct {
// 	Type   AuthTypeEnum   `json:"type"`
// 	Scheme AuthSchemeEnum `json:"scheme"`

// 	Audience  string `json:"aud,omitempty"` // 这个表示 UIAM_ID
// 	Issuer    string `json:"iss,omitempty"`
// 	NotBefore int64  `json:"nbf,omitempty"`
// 	Subject   string `json:"sub,omitempty"`

// 	// Custom
// 	Provider  string `json:"oap"` // Wechat / Dingding
// 	UID       string `json:"uid"` // 这个 UID 表示的是第三方系统的 ID，像 ZOTC 里的那个自己的
// 	SessionID string `json:"sid"`
// 	Sign      string `json:"sig"`
// 	SignAlg   string `json:"sal"`

// 	// ExpriedAt
// 	Duration time.Duration `json:"duration"` //  有效时长
// }
