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

	Type        LoginTypeEnum           `json:"type"`
	Name        string                  `json:"name"`
	PhoneCode   string                  `json:"phone_code"`
	PhoneNumber string                  `json:"phone_number"`
	Email       string                  `json:"email"`
	Password    string                  `json:"password"`
	Captcha     *map[string]interface{} `json:"captcha"`
}

// LoginTypeEnum 枚举
type LoginTypeEnum string

const (
	LoginTypeEnumName    LoginTypeEnum = "name"
	LoginTypeEnumPhone   LoginTypeEnum = "phone"
	LoginTypeEnumEmail   LoginTypeEnum = "email"
	LoginTypeEnumSmsCode LoginTypeEnum = "sms_code"
)

func (e LoginTypeEnum) String() string {
	switch e {
	case LoginTypeEnumName:
		return "name"
	case LoginTypeEnumPhone:
		return "phone"
	case LoginTypeEnumEmail:
		return "email"
	case LoginTypeEnumSmsCode:
		return "sms_code"
	default:
		return ""
	}
}
