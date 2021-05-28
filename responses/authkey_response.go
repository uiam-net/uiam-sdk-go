package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

type AuthorizedKey struct {
	Key         string                   `json:"key"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Scheme      uiammodel.AuthSchemeEnum `json:"scheme"`
	Secret      string                   `json:"secret,omitempty"`
	Scopes      string                   `json:"scopes"`
	Remark      string                   `json:"remark,omitempty"`
	ExpriedAt   *time.Time               `json:"expried_at"`
	CreatedAt   time.Time                `json:"created_at"`
}
