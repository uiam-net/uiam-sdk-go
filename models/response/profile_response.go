package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// ProfileResponse profile
type ProfileResponse struct {
	IdentityID  string                  `json:"identity_id"`
	Name        string                  `json:"name,omitempty"`
	KycLevel    uiammodel.KycLevelEnum  `json:"kyc_level,omitempty"`
	KycStatus   uiammodel.KycStatusEnum `json:"kyc_status,omitempty"`
	KycError    string                  `json:"kyc_error,omitempty"`
	IDDigest    string                  `json:"id_digest,omitempty"`
	CertifiedAt *time.Time              `json:"certified_at"`
}

// NewProfileResponse NewProfileResponse
func NewProfileResponse(profile *uiammodel.Profile) *ProfileResponse {
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
