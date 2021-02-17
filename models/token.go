package uiamsdk

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// JWTPayload RealmJWTPayload
type JWTPayload struct {
	jwtgo.StandardClaims

	UID           string         `json:"uid,omitempty"`
	Mode          string         `json:"mode,omitempty"`
	Scheme        AuthSchemeEnum `json:"scheme,omitempty"`
	OAuthProvider string         `json:"oap,omitempty"` //oap: OAuth Provider
	SessionID     string         `json:"sid,omitempty"`
	Sign          string         `json:"sig,omitempty"`
	SignAlg       string         `json:"sal,omitempty"`
	Extra         string         `json:"extra,omitempty"` // 额外的字段，可以存放 json 等非标数据
}

// Token profile
type Token struct {
	RealmID    string         `json:"realm_id,omitempty"`
	IdentityID string         `json:"identity_id,omitempty"`
	Type       AuthTypeEnum   `json:"type,omitempty"`
	Scheme     AuthSchemeEnum `json:"scheme,omitempty"`
	Provider   string         `json:"oap,omitempty"`
	Key        string         `json:"sid,omitempty"`
	ExpriedAt  *time.Time     `json:"expried_at,omitempty"`
	Token      string         `json:"token,omitempty"`
}
