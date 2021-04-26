package response

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// JWTPayload RealmJWTPayload
type JWTPayloadResponse struct {
	jwtgo.StandardClaims

	UID           string                   `json:"uid,omitempty"`
	Mode          string                   `json:"mode,omitempty"`
	Scheme        uiammodel.AuthSchemeEnum `json:"scheme,omitempty"`
	OAuthProvider string                   `json:"oap,omitempty"` //oap: OAuth Provider
	SessionID     string                   `json:"sid,omitempty"`
	Sign          string                   `json:"sig,omitempty"`
	SignAlg       string                   `json:"sal,omitempty"`
	Extra         string                   `json:"extra,omitempty"` // 额外的字段，可以存放 json 等非标数据
}

// Token profile
type TokenResponse struct {
	RealmID    string                   `json:"realm_id,omitempty"`
	IdentityID string                   `json:"identity_id,omitempty"`
	Type       uiammodel.AuthTypeEnum   `json:"type,omitempty"`
	Scheme     uiammodel.AuthSchemeEnum `json:"scheme,omitempty"`
	Provider   string                   `json:"oap,omitempty"`
	Key        string                   `json:"sid,omitempty"`
	ExpriedAt  *time.Time               `json:"expried_at,omitempty"`
	Token      string                   `json:"token,omitempty"`
}
