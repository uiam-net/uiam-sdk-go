package uiamsdk

import "time"

// Identity Identity
type Identity struct {
	ID             uint64             `mapstructure:"id" json:"id" gorm:"PRIMARY_KEY;NOT NULL"`
	UUID           string             `mapstructure:"uuid" json:"uuid" gorm:"COLUMN:uuid;NOT NULL;TYPE:VARCHAR(36)"`
	Name           string             `mapstructure:"name" json:"name,omitempty" gorm:"COLUMN:name;NOT NULL;TYPE:VARCHAR(36)"`
	Type           IdentityTypeEnum   `mapstructure:"type" json:"type,omitempty" gorm:"COLUMN:type;NOT NULL;TYPE:VARCHAR(18)"`
	RealmUUID      string             `mapstructure:"realm_uuid" json:"realm_uuid,omitempty" gorm:"COLUMN:realm_uuid;NOT NULL;TYPE:VARCHAR(36)"`
	MerchantUUID   string             `mapstructure:"merchant_uuid" json:"merchant_uuid,omitempty" gorm:"COLUMN:merchant_uuid;NOT NULL;TYPE:VARCHAR(36)"`
	PhoneCode      string             `mapstructure:"phone_code" json:"phone_code" gorm:"COLUMN:phone_code;TYPE:VARCHAR(16)"`
	PhoneNumber    string             `mapstructure:"phone_number" json:"phone_number" gorm:"COLUMN:phone_number;TYPE:VARCHAR(32)"`
	PasswordDigest string             `mapstructure:"-" json:"-" gorm:"COLUMN:password_digest;TYPE:VARCHAR(255)"`
	Email          string             `mapstructure:"email" json:"email" gorm:"COLUMN:email;TYPE:VARCHAR(64)"`
	AvatarURL      string             `mapstructure:"avatar_url" json:"avatar_url" gorm:"COLUMN:avatar_url;TYPE:VARCHAR(1024)"`
	Description    string             `mapstructure:"description" json:"description"  gorm:"COLUMN:description;NOT NULL;TYPE:VARCHAR(255)"`
	Remark         string             `mapstructure:"remark" json:"remark" gorm:"COLUMN:remark;NOT NULL;TYPE:VARCHAR(255)"`
	MfaCredential  string             `mapstructure:"mfa_credential" json:"mfa_credential" gorm:"COLUMN:mfa_credential;TYPE:VARCHAR(255)"`
	Attributes     Attribute          `mapstructure:"attributes" json:"attributes" gorm:"COLUMN:attributes;TYPE:json"`
	Status         IdentityStatusEnum `mapstructure:"status" json:"status" gorm:"COLUMN:status;NOT NULL;TYPE:VARCHAR(18)"`
	CreatedAt      time.Time          `mapstructure:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `mapstructure:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time         `json:"-"`
}

// TableName TableName
func (Identity) TableName() string {
	return "identities"
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
