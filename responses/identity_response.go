package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// Identity Identity
type Identity struct {
	IdentityID    uint64                       `json:"-"`
	IdentityUUID  string                       `json:"uuid"`
	Slug          string                       `json:"slug"`
	Username      string                       `json:"username"`
	Nickname      string                       `json:"nickname"`
	PhoneCode     string                       `json:"phone_code"`
	PhoneNumber   string                       `json:"phone_number"`
	PhoneVerified bool                         `json:"phone_verified,omitempty"`
	Email         string                       `json:"email,omitempty"`
	EmailVerified bool                         `json:"email_verified,omitempty"`
	AvatarURL     string                       `json:"avatar_url"`
	Description   string                       `json:"description"`
	MfaType       uiammodel.MfaTypeEnum        `json:"mfa_type,omitempty"`
	MfaStatus     uiammodel.MfaStatusEnum      `json:"mfa_status,omitempty"`
	MfaPayload    uiammodel.Attribute          `json:"mfa_payload,omitempty"`
	Remark        string                       `json:"remark,omitempty"`
	Auths         *Auths                       `json:"authorizations,omitempty"`
	Profile       *Profile                     `json:"profile,omitempty"`
	Attributes    uiammodel.Attribute          `json:"attributes,omitempty"`
	Status        uiammodel.IdentityStatusEnum `json:"status"`
	CreatedAt     time.Time                    `json:"created_at"`
	UpdatedAt     time.Time                    `json:"updated_at"`
}
