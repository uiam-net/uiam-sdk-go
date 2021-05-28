package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// Realm
type Realm struct {
	Slug            string                             `json:"id"`
	Name            string                             `json:"name"`
	Description     string                             `json:"description"`
	AvatarURL       string                             `json:"avatar_url"`
	AuthnModel      uiammodel.RealmAuthzModelEnum      `json:"authn_model"`
	CaptchaMode     uiammodel.RealmCaptchaModeEnum     `json:"captcha_model"`
	CaptchaProvider uiammodel.RealmCaptchaProviderEnum `json:"captcha_provider"`
	CaptchaPayload  uiammodel.Attribute                `json:"captcha_payload"`
	AuthzModel      uiammodel.RealmAuthzModelEnum      `json:"authz_model"`
	PasswordRule    uiammodel.RealmAuthzModelEnum      `json:"password_rule"`
	MfaRule         string                             `json:"mfa_rule"`
	LoginRule       string                             `json:"login_rule"`
	Secret          string                             `json:"secret"`
	SecretCreateAt  *time.Time                         `json:"secret_at"`
	RequestURLs     string                             `json:"request_urls"`
	Attributes      uiammodel.Attribute                `json:"attributes,omitempty"`
	Remark          string                             `json:"remark"`
	Status          uiammodel.RealmStatusEnum          `json:"status"`
	CreatedAt       time.Time                          `json:"created_at"`
}
