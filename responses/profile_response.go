package response

import (
	"time"
)

// Profile
type Profile struct {
	IdentityID  uint64     `json:"identity_id"`
	Name        string     `json:"name,omitempty"`
	CertifiedAt *time.Time `json:"certified_at"`
}
