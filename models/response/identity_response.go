package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// IdentityAuthsResponse Identity
type IdentityAuthsResponse struct {
	IdentityID   string                       `json:"uuid"`
	IdentityName string                       `json:"name"`
	PhoneCode    string                       `json:"phone_code"`
	PhoneNumber  string                       `json:"phone_number"`
	Email        string                       `json:"email"`
	AvatarURL    string                       `json:"avatar_url"`
	Description  string                       `json:"description"`
	Auths        *AuthsResponse               `json:"authorizations,omitempty"`
	Profile      *ProfileResponse             `json:"profile,omitempty"`
	Attributes   uiammodel.Attribute          `json:"attributes"`
	Status       uiammodel.IdentityStatusEnum `json:"status"`
	CreatedAt    time.Time                    `json:"created_at"`
}

// NewIdentityAuthsResponse 构造函数
func NewIdentityAuthsResponse(identity *uiammodel.Identity) *IdentityAuthsResponse {
	response := &IdentityAuthsResponse{
		IdentityID:   identity.UUID,
		IdentityName: identity.Name,
		PhoneCode:    identity.PhoneCode,
		PhoneNumber:  identity.PhoneNumber,
		Email:        identity.Email,
		Description:  identity.Description,
		Status:       identity.Status,
		Attributes:   identity.Attributes,
		CreatedAt:    identity.CreatedAt,
	}

	return response
}

// NewIdentityWithAuthsResponse 构造函数
func NewIdentityWithAuthsResponse(identity *uiammodel.Identity, profile *uiammodel.Profile, auth *uiammodel.Connect) *IdentityAuthsResponse {
	response := &IdentityAuthsResponse{
		IdentityID:   identity.UUID,
		IdentityName: identity.Name,
		PhoneCode:    identity.PhoneCode,
		PhoneNumber:  identity.PhoneNumber,
		Email:        identity.Email,
		Description:  identity.Description,
		Auths:        new(AuthsResponse),
		Attributes:   identity.Attributes,
		Status:       identity.Status,
		CreatedAt:    identity.CreatedAt,
	}

	if profile != nil {
		response.Profile = NewProfileResponse(profile)
	}

	// authsResponse := new(AuthsResponse)
	if auth != nil {
		response.Auths.MixinAuth = NewAuthorizationResponse(auth)
	}

	return response
}
