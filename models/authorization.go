package uiamsdk

import (
	"errors"
	"time"
)

// Authorization Authorization
type Authorization struct {
	ID          uint64               `json:"id" gorm:"PRIMARY_KEY;NOT NULL"`
	IdentityID  string               `json:"identity_id" gorm:"COLUMN:identity_id;NOT NULL;TYPE:BIGINT()"`
	RealmID     string               `json:"realm_id" gorm:"COLUMN:realm_id;NOT NULL;TYPE:VARCHAR(36)"`
	Provider    AuthProviderTypeEnum `json:"provider" gorm:"COLUMN:provider;NOT NULL;TYPE:VARCHAR(36)"`
	OauthID     string               `json:"oauth_id" gorm:"COLUMN:oauth_id;NOT NULL;TYPE:VARCHAR(128)"`
	UnionID     string               `json:"union_id" gorm:"COLUMN:union_id;NOT NULL;TYPE:VARCHAR(128)"`
	AppUserID   string               `json:"app_user_id" gorm:"COLUMN:app_user_id;TYPE:VARCHAR(36)"`
	AppUserName string               `json:"app_user_name" gorm:"COLUMN:app_user_name;TYPE:VARCHAR(36)"`
	Credential  Attribute            `json:"credential" gorm:"COLUMN:credential;TYPE:json"`
	CreatedAt   time.Time            `json:"created_at" gorm:"COLUMN:created_at;NOT NULL"`
	UpdatedAt   time.Time            `json:"updated_at" gorm:"COLUMN:updated_at;NOT NULL"`
}

// TableName TableName
func (Authorization) TableName() string {
	return "authorizations"
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
	// MixinCredentialTypeEnumEdkey edkey
	MixinCredentialTypeEnumEdkey MixinCredentialTypeEnum = "edkey"
	// MixinCredentialTypeEnumToken token
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
