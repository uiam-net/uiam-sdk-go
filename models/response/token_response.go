package response

import (
	"time"

	uiamsdk "github.com/uiam-net/uiam-sdk-go/models"
)

// TokenResponse profile
type TokenResponse struct {
	ReamID    string                          `json:"app_id,omitempty"`
	UserID    string                          `json:"user_id,omitempty"`
	Scheme    uiamsdk.AuthSchemeEnum          `json:"scheme,omitempty"`
	Provider  uiamsdk.ConnectProviderTypeEnum `json:"oap,omitempty"`
	Key       string                          `json:"sid,omitempty"`
	ExpriedAt *time.Time                      `json:"expried_at,omitempty"`
	Token     string                          `json:"token,omitempty"`
}
