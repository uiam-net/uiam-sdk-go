package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// Identity Identity
type Identity struct {
	IdentityID   string                       `json:"uuid"`
	IdentityName string                       `json:"name"`
	PhoneCode    string                       `json:"phone_code"`
	PhoneNumber  string                       `json:"phone_number"`
	Email        string                       `json:"email"`
	AvatarURL    string                       `json:"avatar_url"`
	Description  string                       `json:"description"`
	Auths        *Auths                       `json:"authorizations,omitempty"`
	Profile      *Profile                     `json:"profile,omitempty"`
	Attributes   uiammodel.Attribute          `json:"attributes"`
	Status       uiammodel.IdentityStatusEnum `json:"status"`
	CreatedAt    time.Time                    `json:"created_at"`
}
