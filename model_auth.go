package uiamsdk

import (
	"errors"
	"time"
)

// Authorization Authorization
type Authorization struct {
	ID         uint64               `json:"id"`
	IdentityID string               `json:"identity_id"`
	RealmID    string               `json:"realm_id"`
	Provider   AuthProviderTypeEnum `json:"provider"`
	OauthID    string               `json:"oauth_id"`
	UnionID    string               `json:"union_id"`
	AppUserID  string               `json:"app_user_id"`
	Credential string               `json:"credential"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
}

// AuthorizationList AuthorizationList
type AuthorizationList struct {
	Data       []*Authorization `json:"data"`
	Pagination Pagination       `json:"pagination"`
}

// AuthObject Auths
type AuthObject struct {
	WechatAuth *WechatAuth `json:"wechat,omitempty"`
}

// ================ Wechat ================== //

// WechatCredential WechatCredential
type WechatCredential struct {
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
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

// ================ 枚举 ================== //
// ================ 枚举 ================== //

// AuthProviderTypeEnum 枚举
type AuthProviderTypeEnum string

const (
	// AuthProviderTypeEnumMixin offer
	AuthProviderTypeEnumMixin AuthProviderTypeEnum = "mixin"
	// AuthProviderTypeEnumWechat wechat
	AuthProviderTypeEnumWechat AuthProviderTypeEnum = "wechat"
	// AuthProviderTypeEnumAlipay alipay
	AuthProviderTypeEnumAlipay AuthProviderTypeEnum = "alipay"
)

func (e AuthProviderTypeEnum) String() string {
	switch e {
	case AuthProviderTypeEnumMixin:
		return "mixin"
	case AuthProviderTypeEnumWechat:
		return "wechat"
	case AuthProviderTypeEnumAlipay:
		return "alipay"
	default:
		return ""
	}
}

// NewAuthProviderTypeEnum 构造函数
func NewAuthProviderTypeEnum(provider string) (AuthProviderTypeEnum, error) {
	switch provider {
	case "mixin":
		return AuthProviderTypeEnumMixin, nil
	case "wechat":
		return AuthProviderTypeEnumWechat, nil
	case "alipay":
		return AuthProviderTypeEnumAlipay, nil
	default:
		return "", errors.New("Parse Provider Type Error")
	}
}

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

// AuthSchemeEnum 枚举
type AuthSchemeEnum string

const (
	// AuthSchemeEnumBasic basic
	AuthSchemeEnumBasic AuthSchemeEnum = "basic"
	// AuthSchemeEnumDigest digest
	AuthSchemeEnumDigest AuthSchemeEnum = "digest"
	// AuthSchemeEnumJWTHS jwt_hs
	AuthSchemeEnumJWTHS AuthSchemeEnum = "jwths"
	// AuthSchemeEnumJWTRS jwt_rs
	AuthSchemeEnumJWTRS AuthSchemeEnum = "jwtrs"
	// AuthSchemeEnumJWTES jwt_es
	AuthSchemeEnumJWTES AuthSchemeEnum = "jwtes"
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
	default:
		return ""
	}
}
