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
	RealmID string       `json:"realm,omitempty"`
	Captcha *CaptchaData `json:"captcha,omitempty"`
}

type LoginResp interface {
	GetType() uiammodel.SessionTypeEnum
}

// LoginCaptcha 登录请求后返回类型  LoginCaptcha
type LoginCaptcha struct {
	Type           uiammodel.SessionTypeEnum          `json:"type,omitempty"`
	RealmID        string                             `json:"realm_id,omitempty"`
	Mode           uiammodel.RealmCaptchaModeEnum     `json:"mode,omitempty"`
	Provider       uiammodel.RealmCaptchaProviderEnum `json:"provider"`
	CaptchaPayload interface{}                        `json:"payload"`
}

func (captcha LoginCaptcha) GetType() uiammodel.SessionTypeEnum {
	return captcha.Type
}

// LoginMfa 登录请求后返回类型 LoginMfa
type LoginMfa struct {
	Type       uiammodel.SessionTypeEnum `json:"type,omitempty"`
	RealmID    string                    `json:"realm_id,omitempty"`
	Mode       uiammodel.MfaTypeEnum     `json:"mode"` // otp/sms
	TempID     string                    `json:"temp_id,omitempty"`
	MfaPayload interface{}               `json:"payload,omitempty"`
}

func (mfa LoginMfa) GetType() uiammodel.SessionTypeEnum {
	return mfa.Type
}

// LoginToken 登录请求后返回类型 LoginToken
type LoginToken struct {
	Type         uiammodel.SessionTypeEnum `json:"type,omitempty"`
	RealmID      string                    `json:"realm_id,omitempty"`
	IdentityUUID string                    `json:"identity_id,omitempty"`
	ExpriedAt    *time.Time                `json:"expried_at,omitempty"`
	Token        string                    `json:"token,omitempty"`
}

func (token LoginToken) GetType() uiammodel.SessionTypeEnum {
	return token.Type
}
