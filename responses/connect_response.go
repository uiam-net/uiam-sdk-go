package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// Connect Identity
type Connect struct {
	ID           uint64                        `json:"id"`
	IdentityUUID string                        `json:"identity_uuid"`
	RealmUUID    string                        `json:"realm_uuid"`
	Provider     uiammodel.ConnectProviderEnum `json:"provider"`
	OauthID      string                        `json:"oauth_id"`
	UnionID      string                        `json:"union_id"`
	AppUserID    string                        `json:"app_user_id"`
	AppUserName  string                        `json:"app_user_name"`
	Credential   uiammodel.Attribute           `json:"credential"`
	CreatedAt    time.Time                     `json:"created_at"`
}
