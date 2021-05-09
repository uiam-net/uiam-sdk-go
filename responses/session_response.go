package response

import (
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

type RealmCaptchaTypeDigtal struct {
	Image string `json:"image,omitempty"`
}

type LoginData struct {
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
	Provider uiammodel.RealmCaptchaProviderEnum `json:"provider"`
	Payload  interface{}                        `json:"payload"`
}
