package uiamsdk

import "time"

// User User
type User struct {
	UUID        string      `mapstructure:"uuid" json:"uuid"`
	Name        string      `mapstructure:"name" json:"name,omitempty"`
	PhoneCode   string      `mapstructure:"phone_code" json:"phone_code"`
	PhoneNumber string      `mapstructure:"phone_number" json:"phone_number"`
	Email       string      `mapstructure:"email" json:"email"`
	Description string      `mapstructure:"description" json:"description"`
	Attributes  Attribute   `mapstructure:"attributes" json:"attributes"`
	Auths       *AuthObject `mapstructure:"authorizations" json:"authorizations,omitempty"`
	Profile     *Profile    `mapstructure:"profile" json:"profile,omitempty"`
	Status      string      `mapstructure:"status" json:"status"`
	CreatedAt   time.Time   `mapstructure:"created_at" json:"created_at"`
	AvatarURL   string      `mapstructure:"avatar_url" json:"avatar_url"`
}

// Profile profile
type Profile struct {
	UserID      string    `json:"user_id"`
	Name        string    `json:"name,omitempty"`
	KycLevel    int       `json:"kyc_level,omitempty"`
	KycStatus   string    `json:"kyc_status,omitempty"`
	KycError    string    `json:"kyc_error,omitempty"`
	IDDigest    string    `json:"id_digest,omitempty"`
	CertifiedAt time.Time `json:"certified_at"`
}

// =============== IdentityTypeEnum =============== //

// IdentityTypeEnum 枚举
type IdentityTypeEnum string

const (
	// IdentityTypeEnumStd 正常注册用户
	IdentityTypeEnumStd IdentityTypeEnum = "std"
	// IdentityTypeEnumSys 系统调用使用
	IdentityTypeEnumSys IdentityTypeEnum = "sys"
	// IdentityTypeEnumVir 系统需要注册的虚拟用户
	IdentityTypeEnumVir IdentityTypeEnum = "vir"
)

func (e IdentityTypeEnum) String() string {
	switch e {
	case IdentityTypeEnumStd:
		return "std"
	case IdentityTypeEnumSys:
		return "sys"
	case IdentityTypeEnumVir:
		return "vir"
	default:
		return "unkonwn"
	}
}

// IdentityStatusEnum 枚举
type IdentityStatusEnum string

const (
	// IdentityStatusEnumInit 禁用
	IdentityStatusEnumInit IdentityStatusEnum = "init"
	// IdentityStatusEnumPending 正常
	IdentityStatusEnumPending IdentityStatusEnum = "pending"
	// IdentityStatusEnumNormal 正常
	IdentityStatusEnumNormal IdentityStatusEnum = "normal"
	// IdentityStatusEnumDisabled 禁用
	IdentityStatusEnumDisabled IdentityStatusEnum = "disabled"
)

func (e IdentityStatusEnum) String() string {
	switch e {
	case IdentityStatusEnumInit:
		return "init"
	case IdentityStatusEnumPending:
		return "pending"
	case IdentityStatusEnumNormal:
		return "normal"
	case IdentityStatusEnumDisabled:
		return "disabled"
	default:
		return "unkonwn"
	}
}
