package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// Profile
type Profile struct {
	IdentityID  string                  `json:"identity_id"`
	Name        string                  `json:"name,omitempty"`
	KycLevel    uiammodel.KycLevelEnum  `json:"kyc_level,omitempty"`
	KycStatus   uiammodel.KycStatusEnum `json:"kyc_status,omitempty"`
	KycError    string                  `json:"kyc_error,omitempty"`
	IDDigest    string                  `json:"id_digest,omitempty"`
	CertifiedAt *time.Time              `json:"certified_at"`
}
