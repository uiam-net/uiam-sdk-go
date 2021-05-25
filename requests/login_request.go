package uiamsdk

import (
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

type CaptchaUiamDigtal struct {
	Id   string `json:"captcha_id"`
	Data string `json:"data"`
}

// CaptchaData
type CaptchaData struct {
	Mode     uiammodel.RealmCaptchaModeEnum     `json:"mode"`
	Provider uiammodel.RealmCaptchaProviderEnum `json:"provider"`
	Payload  map[string]interface{}             `json:"payload"`
}

// LoginRequest LoginRequest
type LoginRequest struct {
	BaseRequest

	Type        uiammodel.LoginTypeEnum `json:"type"`
	Username    string                  `json:"username"`
	PhoneCode   string                  `json:"phone_code"`
	PhoneNumber string                  `json:"phone_number"`
	Email       string                  `json:"email"`
	Password    string                  `json:"password"`
	Captcha     *map[string]interface{} `json:"captcha"`
}
