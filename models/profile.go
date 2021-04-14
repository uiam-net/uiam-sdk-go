package uiamsdk

import (
	"time"
)

// Profile profile
type Profile struct {
	ID                uint64        `json:"id" gorm:"PRIMARY_KEY;NOT NULL"`
	IdentityID        string        `json:"identity_uuid" gorm:"COLUMN:identity_uuid;NOT NULL;TYPE:BIGINT(20)"`
	Country           CountryEnum   `json:"country" gorm:"COLUMN:country;NOT NULL;TYPE:VARCHAR(18)"`
	Name              string        `json:"name" gorm:"COLUMN:name;NOT NULL;TYPE:VARCHAR(18)"`
	Gender            GenderEnum    `json:"gender" gorm:"COLUMN:gender;NOT NULL;TYPE:VARCHAR(2)"`
	Nation            string        `json:"nation" gorm:"COLUMN:nation;NOT NULL;TYPE:VARCHAR(36)"`
	IDType            IDTypeEnum    `json:"id_type" gorm:"COLUMN:id_type;NOT NULL;TYPE:INT(11)"`
	IDNoEncrypted     string        `json:"id_no_encrypted" gorm:"COLUMN:id_no_encrypted;NOT NULL;TYPE:VARCHAR(255)"`
	BirthEncrypted    string        `json:"birth_encrypted" gorm:"COLUMN:birth_encrypted;NOT NULL;TYPE:VARCHAR(255)"`
	IDDigest          string        `json:"id_digest" gorm:"COLUMN:id_digest;NOT NULL;TYPE:VARCHAR(255)"`
	IssuedByEncrypted string        `json:"issued_by_encrypted" gorm:"COLUMN:issued_by_encrypted;NOT NULL;TYPE:VARCHAR(512)"`
	AddressEncrypted  string        `json:"address_encrypted" gorm:"COLUMN:address_encrypted;NOT NULL;TYPE:VARCHAR(1024)"`
	KycLevel          KycLevelEnum  `json:"kyc_level" gorm:"COLUMN:kyc_level;NOT NULL"`
	KycStatus         KycStatusEnum `json:"kyc_status" gorm:"COLUMN:kyc_status;NOT NULL;TYPE:VARCHAR(18)"`
	KycError          string        `json:"kyc_error" gorm:"COLUMN:kyc_error;NOT NULL;TYPE:VARCHAR(64)"`
	EffectedAt        string        `json:"effected_at" gorm:"COLUMN:effected_at;NOT NULL;TYPE:VARCHAR(10)"`
	ExpiredAt         string        `json:"expired_at" gorm:"COLUMN:expired_at;NOT NULL;TYPE:VARCHAR(10)"`
	CertifiedAt       *time.Time    `json:"certified_at" gorm:"COLUMN:certified_at;"`
	CreatedAt         *time.Time    `json:"created_at" gorm:"COLUMN:created_at;NOT NULL"`
	UpdatedAt         *time.Time    `json:"updated_at" gorm:"COLUMN:updated_at;NOT NULL"`
}

// =============== KycStatusEnum =============== //

// KycStatusEnum 枚举
type KycStatusEnum string

const (
	// KycStatusEnumProcessing 处理中
	KycStatusEnumProcessing KycStatusEnum = "processing"
	// KycStatusEnumRejected rejected
	KycStatusEnumRejected KycStatusEnum = "rejected"
	// KycStatusEnumApproved approved
	KycStatusEnumApproved KycStatusEnum = "approved"
	// KycStatusEnumUnkonwn unkonwn
	KycStatusEnumUnkonwn KycStatusEnum = "unkonwn"
)

func (e KycStatusEnum) String() string {
	switch e {
	case KycStatusEnumProcessing:
		return "processing"
	case KycStatusEnumRejected:
		return "rejected"
	case KycStatusEnumApproved:
		return "approved"
	default:
		return "unkonwn"
	}
}

// NewKycStatusEnum NewKycStatusEnum
func NewKycStatusEnum(enum string) KycStatusEnum {
	switch enum {
	case "processing":
		return KycStatusEnumProcessing
	case "rejected":
		return KycStatusEnumRejected
	case "approved":
		return KycStatusEnumApproved
	default:
		return KycStatusEnumUnkonwn
	}
}
