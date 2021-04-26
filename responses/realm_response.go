package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// Realm
type Realm struct {
	UUID         string                    `json:"uuid"`
	Slug         string                    `json:"slug"`
	MerchantUUID string                    `json:"merchant_uuid"`
	Name         string                    `json:"name"`
	AuthModel    string                    `json:"auth_model"`
	Description  string                    `json:"description"`
	Remark       string                    `json:"remark"`
	Status       uiammodel.RealmStatusEnum `json:"status"`
	CreatedAt    time.Time                 `json:"created_at"`
}
