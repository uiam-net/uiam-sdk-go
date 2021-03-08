package uiamsdk

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// TokenCreateRequest TokenCreateRequest
type TokenCreateRequest struct {
	BaseRequest

	Audience  string        `json:"aud,omitempty"` // IdentityID
	Issuer    string        `json:"iss,omitempty"` // RealmID
	Subject   string        `json:"sub,omitempty"` // Subject
	NotBefore int64         `json:"nbf,omitempty"` //
	Duration  time.Duration `json:"duration"`      // ExpriedAt 有效时长

	// Custom
	Type      uiammodel.AuthTypeEnum         `json:"type"`
	Provider  uiammodel.AuthProviderTypeEnum `json:"oap,omitempty"`
	Scheme    uiammodel.AuthSchemeEnum       `json:"scheme,omitempty"`
	UID       string                         `json:"uid,omitempty"` // Audience
	SessionID string                         `json:"sid,omitempty"` // AuthKey
	Sign      string                         `json:"sig,omitempty"` // Sign
	SignAlg   string                         `json:"sal,omitempty"` // SignAlg
	Extra     string                         `json:"extra,omitempty"`
}
