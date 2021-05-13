package uiamsdk

import (
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// RealmCreateRequest
type RealmCreateRequest struct {
	BaseRequest

	Name        string `json:"name"`
	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

// CaptchaPayloadUiam Realm
type CaptchaPayloadUiam struct {
	Type   uiammodel.RealmCaptchaTypeEnum `json:"type,omitempty"`
	Length uint32                         `json:"length,omitempty"`
	Width  uint32                         `json:"width,omitempty"`
	Height uint32                         `json:"height,omitempty"`
}
