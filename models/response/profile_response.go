package response

import (
	"time"

	uiammodels "github.com/uiam-net/uiam-sdk-go/models"
)

// ProfileResponse profile
type ProfileResponse struct {
	IdentityID  string                   `json:"identity_id"`
	Name        string                   `json:"name,omitempty"`
	KycLevel    uiammodels.KycLevelEnum  `json:"kyc_level,omitempty"`
	KycStatus   uiammodels.KycStatusEnum `json:"kyc_status,omitempty"`
	KycError    string                   `json:"kyc_error,omitempty"`
	IDDigest    string                   `json:"id_digest,omitempty"`
	CertifiedAt *time.Time               `json:"certified_at"`
}

// NewProfileResponse NewProfileResponse
func NewProfileResponse(profile *uiammodels.Profile) *ProfileResponse {
	response := &ProfileResponse{
		IdentityID:  profile.IdentityID,
		Name:        profile.Name,
		KycLevel:    profile.KycLevel,
		KycStatus:   profile.KycStatus,
		KycError:    profile.KycError,
		CertifiedAt: profile.CertifiedAt,
		IDDigest:    profile.IDDigest,
	}

	return response
}
