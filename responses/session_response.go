package response

import (
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

type RealmCaptchaTypeDigtal struct {
	Image string `json:"image,omitempty"`
}

type LoginData struct {
	Captcha *Captcha `json:"captcha,omitempty"`
}

// JWTPayload
type Captcha struct {
	Mode uiammodel.RealmCaptchaModeEnum `json:"mode,omitempty"`
	Type uiammodel.RealmCaptchaTypeEnum `json:"type,omitempty"`
	Data uiammodel.Attribute            `json:"data,omitempty"`
}
