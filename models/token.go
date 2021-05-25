package uiamsdk

// JWTPayload
type JWTPayload struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`

	UID           string         `json:"uid,omitempty"`
	RealmID       string         `json:"realm_id,omitempty"`
	RealmSlug     string         `json:"realm,omitempty"`
	MerchantSlug  string         `json:"merchant,omitempty"`
	Mode          string         `json:"mode,omitempty"`
	Scheme        AuthSchemeEnum `json:"scheme,omitempty"`
	OAuthProvider string         `json:"oap,omitempty"` //oap: OAuth Provider
	SessionID     string         `json:"sid,omitempty"`
	Sign          string         `json:"sig,omitempty"`
	SignAlg       string         `json:"sal,omitempty"`
	Extra         string         `json:"extra,omitempty"` // 额外的字段，可以存放 json 等非标数据
}

func (c JWTPayload) Valid() error {
	return nil
}
