package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

type Merchant struct {
	ID          uint64                       `json:"-"`
	Slug        string                       `json:"slug"`
	Type        uiammodel.MerchantTypeEnum   `mapstructure:"type" json:"type,omitempty"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Remark      string                       `json:"remark"`
	Status      uiammodel.MerchantStatusEnum `json:"status"`
	CreatedAt   time.Time                    `json:"created_at"`
}
