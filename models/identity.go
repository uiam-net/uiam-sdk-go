package uiamsdk

// =============== IdentityTypeEnum =============== //

// IdentityTypeEnum 枚举
type IdentityTypeEnum string

const (
	IdentityTypeEnumStd   IdentityTypeEnum = "std"
	IdentityTypeEnumAdmin IdentityTypeEnum = "admin"
	IdentityTypeEnumSys   IdentityTypeEnum = "sys"
)

func (e IdentityTypeEnum) String() string {
	switch e {
	case IdentityTypeEnumStd:
		return "std"
	case IdentityTypeEnumAdmin:
		return "admin"
	case IdentityTypeEnumSys:
		return "sys"
	default:
		return ""
	}
}

// IdentityStatusEnum 枚举
type IdentityStatusEnum string

const (
	IdentityStatusEnumInit     IdentityStatusEnum = "init"
	IdentityStatusEnumPending  IdentityStatusEnum = "pending"
	IdentityStatusEnumNormal   IdentityStatusEnum = "normal"
	IdentityStatusEnumLocked   IdentityStatusEnum = "locked"
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
		return ""
	}
}

// MfaTypeEnum 枚举
type MfaTypeEnum string

const (
	MfaTypeEnumNone        MfaTypeEnum = "none"
	MfaTypeEnumOtp         MfaTypeEnum = "otp"
	MfaTypeEnumSmsCode     MfaTypeEnum = "sms_code"
	MfaTypeEnumEmailCode   MfaTypeEnum = "email_code"
	MfaTypeEnumFaceprint   MfaTypeEnum = "faceprint"   // 面纹
	MfaTypeEnumFingerprint MfaTypeEnum = "fingerprint" // 指纹
)

func (e MfaTypeEnum) String() string {
	switch e {
	case MfaTypeEnumOtp:
		return "otp"
	case MfaTypeEnumSmsCode:
		return "sms_code"
	case MfaTypeEnumEmailCode:
		return "email_code"
	case MfaTypeEnumFaceprint:
		return "faceprint"
	case MfaTypeEnumFingerprint:
		return "fingerprint"
	default:
		return "none"
	}
}

// MfaTypeEnum 枚举
type MfaStatusEnum string

const (
	MfaStatusEnumActive    MfaStatusEnum = "active"
	MfaStatusEnumDisactive MfaStatusEnum = "disactive"
)

func (e MfaStatusEnum) String() string {
	switch e {
	case MfaStatusEnumActive:
		return "active"
	case MfaStatusEnumDisactive:
		return "disactive"
	default:
		return "disactive"
	}
}
