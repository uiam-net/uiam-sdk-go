package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// Realm
type Realm struct {
	Slug        string                    `json:"slug"`
	Name        string                    `json:"name"`
	AvatarURL   string                    `json:"avatar_url"`
	AuthModel   string                    `json:"auth_model"`
	Description string                    `json:"description"`
	Remark      string                    `json:"remark"`
	Status      uiammodel.RealmStatusEnum `json:"status"`
	CreatedAt   time.Time                 `json:"created_at"`
}
