package response

import (
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

type MfaOtpPayload struct {
	URL         string `json:"url"`
	QRCodeImage string `json:"image,omitempty"`
}

type Mfa struct {
	IdentityID string                `json:"identity_id"`
	Mode       uiammodel.MfaTypeEnum `json:"mode,omitempty"`
	Payload    interface{}           `json:"payload,omitempty"`
}
