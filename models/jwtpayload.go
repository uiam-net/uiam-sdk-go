package uiamsdk

import (
	jwtgo "github.com/dgrijalva/jwt-go"
)

// AppJWTPayload AppJWTPayload
type AppJWTPayload struct {
	jwtgo.StandardClaims

	UID           string         `json:"uid,omitempty"`
	Scheme        AuthSchemeEnum `json:"scheme,omitempty"`
	OAuthProvider string         `json:"oap,omitempty"` //oap: OAuth Provider
	SessionID     string         `json:"sid,omitempty"`
	Sign          string         `json:"sig,omitempty"`
	SignAlg       string         `json:"sal,omitempty"`
	Extra         string         `json:"extra,omitempty"` // 额外的字段，可以存放 json 等非标数据
}
