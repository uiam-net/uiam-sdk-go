package uiamsdk

import (
	"errors"
	"time"
)

// ================ 枚举 ================== //
// ================ 枚举 ================== //

// ConnectProviderEnum 枚举
type ConnectProviderEnum string

const (
	ConnectProviderEnumGoogle    ConnectProviderEnum = "google"
	ConnectProviderEnumMicrosoft ConnectProviderEnum = "microsoft"
	ConnectProviderEnumFacebook  ConnectProviderEnum = "facebook"
	ConnectProviderEnumTwitter   ConnectProviderEnum = "twitter"
	ConnectProviderEnumWechat    ConnectProviderEnum = "wechat"
	ConnectProviderEnumQQ        ConnectProviderEnum = "qq"
	ConnectProviderEnumAlipay    ConnectProviderEnum = "alipay"
	ConnectProviderEnumDingtalk  ConnectProviderEnum = "dingtalk"
	ConnectProviderEnumWeibo     ConnectProviderEnum = "weibo"
	ConnectProviderEnumLADP      ConnectProviderEnum = "ladp"
)

func (e ConnectProviderEnum) String() string {
	switch e {

	case ConnectProviderEnumGoogle:
		return "google"
	case ConnectProviderEnumMicrosoft:
		return "microsoft"
	case ConnectProviderEnumFacebook:
		return "facebook"
	case ConnectProviderEnumTwitter:
		return "twitter"
	case ConnectProviderEnumWechat:
		return "wechat"
	case ConnectProviderEnumQQ:
		return "qq"
	case ConnectProviderEnumAlipay:
		return "alipay"
	case ConnectProviderEnumDingtalk:
		return "dingtalk"
	case ConnectProviderEnumWeibo:
		return "weibo"
	case ConnectProviderEnumLADP:
		return "ladp"
	default:
		return ""
	}
}

// NewConnectProviderEnum 构造函数
func NewConnectProviderEnum(provider string) (ConnectProviderEnum, error) {
	switch provider {
	case "wechat":
		return ConnectProviderEnumWechat, nil
	case "alipay":
		return ConnectProviderEnumAlipay, nil
	default:
		return "", errors.New("Parse Provider Type Error")
	}
}

// ========================= AuthTypeEnum ========================= //

// AuthTypeEnum 枚举
type AuthTypeEnum string

const (
	// AuthTypeEnumUser user
	AuthTypeEnumUser AuthTypeEnum = "user"
	// AuthTypeEnumApp app
	AuthTypeEnumApp AuthTypeEnum = "app"
	// AuthTypeEnumAppuser appuser
	AuthTypeEnumAppuser AuthTypeEnum = "appuser"
	// AuthTypeEnumSystem system
	AuthTypeEnumSystem AuthTypeEnum = "system"
)

func (e AuthTypeEnum) String() string {
	switch e {
	case AuthTypeEnumUser:
		return "user"
	case AuthTypeEnumApp:
		return "app"
	case AuthTypeEnumAppuser:
		return "appuser"
	case AuthTypeEnumSystem:
		return "system"
	default:
		return ""
	}
}

// ========================= AuthSchemeEnum ========================= //

// AuthSchemeEnum 枚举
type AuthSchemeEnum string

const (
	AuthSchemeEnumBasic  AuthSchemeEnum = "basic"
	AuthSchemeEnumDigest AuthSchemeEnum = "digest"
	AuthSchemeEnumJWTHS  AuthSchemeEnum = "jwths"
	AuthSchemeEnumJWTRS  AuthSchemeEnum = "jwtrs"
	AuthSchemeEnumJWTES  AuthSchemeEnum = "jwtes"
	AuthSchemeEnumSAML   AuthSchemeEnum = "saml"
	AuthSchemeEnumCAS    AuthSchemeEnum = "cas"
	AuthSchemeEnumOIDC   AuthSchemeEnum = "oidc"
	AuthSchemeEnumLDAP   AuthSchemeEnum = "ldap"
)

func (e AuthSchemeEnum) String() string {
	switch e {
	case AuthSchemeEnumBasic:
		return "basic"
	case AuthSchemeEnumDigest:
		return "digest"
	case AuthSchemeEnumJWTHS:
		return "jwths"
	case AuthSchemeEnumJWTRS:
		return "jwtrs"
	case AuthSchemeEnumJWTES:
		return "jwtes"
	case AuthSchemeEnumSAML:
		return "saml"
	case AuthSchemeEnumCAS:
		return "cas"
	case AuthSchemeEnumOIDC:
		return "oidc"
	case AuthSchemeEnumLDAP:
		return "ldap"
	default:
		return ""
	}
}

// ==================== Credential ================== //
// ==================== Credential ================== //

// ================ Wechat ================== //

// WechatCredential WechatCredential
type WechatCredential struct {
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	SessionKey   string `json:"session_key,omitempty"` // Weapp 加密数据用
}

// WechatAuth WechatAuth
type WechatAuth struct {
	IdentityID string            `json:"identity_id"`
	RealmID    string            `json:"realm_id"`
	Provider   string            `json:"provider"`
	OauthID    string            `json:"oauth_id"`
	UnionID    string            `json:"union_id"`
	UserName   string            `json:"user_name"`
	Credential *WechatCredential `json:"credential"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

// ================ Mixin ================== //

// MixinCredential MixinCredential
type MixinCredential struct {
	Type           MixinCredentialTypeEnum `json:"type"`
	EdPrivKey      string                  `json:"ed_priv_key"`       // edkey
	EdServerPubKey string                  `json:"ed_server_pub_key"` // edkey
	ClientID       string                  `json:"client_id"`         // edkey
	AuthID         string                  `json:"auth_id"`           // edkey
	Scope          string                  `json:"scope"`             // token & edkey
	AccessToken    string                  `json:"access_token"`      // token
}

// MixinCredentialTypeEnum 枚举
type MixinCredentialTypeEnum string

const (
	MixinCredentialTypeEnumEdkey MixinCredentialTypeEnum = "edkey"
	MixinCredentialTypeEnumToken MixinCredentialTypeEnum = "token"
)

func (e MixinCredentialTypeEnum) String() string {
	switch e {
	case MixinCredentialTypeEnumEdkey:
		return "edkey"
	case MixinCredentialTypeEnumToken:
		return "token"
	default:
		return ""
	}
}
