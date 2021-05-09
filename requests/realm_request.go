package uiamsdk

import (
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// CaptchaPayloadUiam Realm
type CaptchaPayloadUiam struct {
	Type   uiammodel.RealmCaptchaTypeEnum `json:"type,omitempty"`
	Length uint32                         `json:"length,omitempty"`
	Width  uint32                         `json:"width,omitempty"`
	Height uint32                         `json:"height,omitempty"`
}
