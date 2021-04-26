package response

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// AuthsResponse AuthsResponse
type AuthsResponse struct {
	MixinAuth  interface{} `json:"mixin,omitempty"`
	WechatAuth interface{} `json:"wechat,omitempty"`
}

// AuthVerifyResponse User
type AuthVerifyResponse struct {
	App       string    `json:"app"`
	User      string    `json:"user"`
	Status    string    `json:"status"`
	Error     string    `json:"error"`
	CreatedAt time.Time `json:"created_at"`
}

// AuthorizationResponse AuthorizationResponse
type AuthorizationResponse struct {
	IdentityID  string                            `json:"identity_id"`
	RealmID     string                            `json:"realm_id"`
	Provider    uiammodel.ConnectProviderTypeEnum `json:"provider"`
	OauthID     string                            `json:"oauth_id"`
	UnionID     string                            `json:"union_id"`
	AppUserID   string                            `json:"app_user_id"`
	AppUserName string                            `json:"app_user_name"`
	Credential  uiammodel.Attribute               `json:"credential"`
	CreatedAt   time.Time                         `json:"created_at"`
	UpdatedAt   time.Time                         `json:"updated_at"`
}

// ======================== mixin ============================ //

// MixinCredentialResp MixinCredential
type MixinCredentialResp struct {
	Type           uiammodel.MixinCredentialTypeEnum `json:"type,omitempty"`
	EdPrivKey      string                            `json:"ed_priv_key,omitempty"`       // edkey
	EdServerPubKey string                            `json:"ed_server_pub_key,omitempty"` // edkey
	ClientID       string                            `json:"client_id,omitempty"`         // edkey
	AuthID         string                            `json:"auth_id,omitempty"`           // edkey
	Scope          string                            `json:"scope,omitempty"`             // token & edkey
	AccessToken    string                            `json:"access_token,omitempty"`      // token
}

// MixinAuthCredentialResponse MixinAuthCredential
type MixinAuthCredentialResponse struct {
	IdentityID string                            `json:"identity_id"`
	RealmID    string                            `json:"app_id"`
	Provider   uiammodel.ConnectProviderTypeEnum `json:"provider"`
	OauthID    string                            `json:"oauth_id"`
	MixinID    string                            `json:"mixin_id"`
	UserName   string                            `json:"user_name"`
	Credential *MixinCredentialResp              `json:"credential"`
	CreatedAt  time.Time                         `json:"created_at"`
	UpdatedAt  time.Time                         `json:"updated_at"`
}
