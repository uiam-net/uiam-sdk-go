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

// NewAuthorizationResponse NewAuthorizationResponse
func NewAuthorizationResponse(auth *uiammodel.Connect) *AuthorizationResponse {
	authResp := new(AuthorizationResponse)

	authResp.IdentityID = auth.IdentityUUID
	authResp.RealmID = auth.RealmUUID
	authResp.Provider = auth.Provider
	authResp.OauthID = auth.OauthID
	authResp.UnionID = auth.UnionID
	authResp.AppUserID = auth.AppUserID
	authResp.AppUserName = auth.AppUserName
	authResp.Credential = auth.Credential
	authResp.CreatedAt = auth.CreatedAt
	authResp.UpdatedAt = auth.UpdatedAt

	return authResp
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

// NewMixinAuthCredentialResp NewMixinAuthCredential
func NewMixinAuthCredentialResp(auth *uiammodel.Connect) *MixinAuthCredentialResponse {
	// Deserialization Credential
	mixinCredential := new(MixinCredentialResp)

	// if auth.Credential != "" {
	// 	var mapCredential map[string]interface{}

	// 	if err := json.Unmarshal([]byte(auth.Credential), &mapCredential); err == nil {
	// 		switch mapCredential["type"] {
	// 		case uiammodel.MixinCredentialTypeEnumEdkey.String():
	// 			json.Unmarshal([]byte(auth.Credential), mixinCredential)
	// 			mixinCredential.Type = uiammodel.MixinCredentialTypeEnumEdkey
	// 		case uiammodel.MixinCredentialTypeEnumToken.String():
	// 			json.Unmarshal([]byte(auth.Credential), mixinCredential)
	// 			mixinCredential.Type = uiammodel.MixinCredentialTypeEnumToken
	// 		default:
	// 			json.Unmarshal([]byte(auth.Credential), mixinCredential)
	// 			mixinCredential.Type = uiammodel.MixinCredentialTypeEnumToken
	// 		}
	// 	}
	// }

	mixinAuth := new(MixinAuthCredentialResponse)
	mixinAuth.IdentityID = auth.IdentityUUID
	mixinAuth.RealmID = auth.RealmUUID
	mixinAuth.Provider = auth.Provider
	mixinAuth.OauthID = auth.OauthID
	mixinAuth.MixinID = auth.AppUserID
	mixinAuth.UserName = auth.AppUserName
	mixinAuth.Credential = mixinCredential
	mixinAuth.CreatedAt = auth.CreatedAt
	mixinAuth.UpdatedAt = auth.UpdatedAt

	return mixinAuth
}
