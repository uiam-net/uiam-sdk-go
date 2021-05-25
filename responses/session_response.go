package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// LoginConfig 登录前返回 LoginConfig
type LoginMfaConfig struct {
	RealmID string       `json:"realm,omitempty"`
	Captcha *CaptchaData `json:"captcha,omitempty"`
}

type CaptchaUiamDigtal struct {
	Length    int    `json:"length,omitempty"`
	Height    int    `json:"-"`
	Width     int    `json:"-"`
	CaptchaID string `json:"captcha_id,omitempty"`
	Data      string `json:"data,omitempty"`
}

// CaptchaData
type CaptchaData struct {
	Mode     uiammodel.RealmCaptchaModeEnum     `json:"mode"`
	Provider uiammodel.RealmCaptchaProviderEnum `json:"provider,omitempty"`
	Payload  interface{}                        `json:"payload,omitempty"`
}

// LoginConfig 登录前返回 LoginConfig
type LoginConfig struct {
	RealmSlug string       `json:"realm_slug,omitempty"`
	Captcha   *CaptchaData `json:"captcha,omitempty"`
}

type LoginResp struct {
	Type    uiammodel.SessionTypeEnum `json:"type,omitempty"`
	Captcha *LoginCaptcha
	Mfa     *LoginMfa
	Token   *LoginToken
}

// LoginCaptcha 登录请求后返回类型  LoginCaptcha
type LoginCaptcha struct {
	Type           uiammodel.SessionTypeEnum          `json:"type,omitempty"`
	RealmSlug      string                             `json:"realm_slug,omitempty"`
	Mode           uiammodel.RealmCaptchaModeEnum     `json:"mode,omitempty"`
	Provider       uiammodel.RealmCaptchaProviderEnum `json:"provider"`
	CaptchaPayload interface{}                        `json:"payload"`
}

// LoginMfa 登录请求后返回类型 LoginMfa
type LoginMfa struct {
	Type       uiammodel.SessionTypeEnum `json:"type,omitempty"`
	RealmSlug  string                    `json:"realm_slug,omitempty"`
	Mode       uiammodel.MfaTypeEnum     `json:"mode"` // otp/sms
	TempID     string                    `json:"temp_id,omitempty"`
	MfaPayload interface{}               `json:"payload,omitempty"`
}

// LoginToken 登录请求后返回类型 LoginToken
type LoginToken struct {
	Type         uiammodel.SessionTypeEnum `json:"type,omitempty"`
	RealmSlug    string                    `json:"realm_slug,omitempty"`
	IdentityUUID string                    `json:"identity_id,omitempty"`
	ExpriedAt    *time.Time                `json:"expried_at,omitempty"`
	Token        string                    `json:"token,omitempty"`
}
