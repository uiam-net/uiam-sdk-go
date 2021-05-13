package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

type AuthorizedKey struct {
	IdentityUUID string                   `json:"identity_uuid"`
	ReamSlug     string                   `json:"realm_uuid"`
	Name         string                   `json:"name"`
	Description  string                   `json:"description"`
	Scheme       uiammodel.AuthSchemeEnum `json:"scheme"`
	Scopes       string                   `json:"scopes"`
	Key          string                   `json:"auth_key"`
	Secret       string                   `json:"secret"`
	Remark       string                   `json:"remark"`
	ExpriedAt    *time.Time               `json:"expried_at"`
	CreatedAt    time.Time                `json:"created_at"`
}
