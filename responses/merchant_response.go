package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

type MerchantResponse struct {
	ID          uint32                       `json:"id"`
	UUID        string                       `mapstructure:"uuid" json:"uuid"`
	Type        uiammodel.MerchantTypeEnum   `mapstructure:"type" json:"type,omitempty"`
	Slug        string                       `json:"slug"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Remark      string                       `json:"remark"`
	Status      uiammodel.MerchantStatusEnum `json:"status"`
	CreatedAt   time.Time                    `json:"created_at"`
}
