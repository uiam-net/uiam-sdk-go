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

// MfaModeEnum 枚举
type MfaModeEnum string

const (
	MfaModeEnumNone MfaModeEnum = "none"
	MfaModeEnumOtp  MfaModeEnum = "otp"
	MfaModeEnumSms  MfaModeEnum = "sms"
)

func (e MfaModeEnum) String() string {
	switch e {
	case MfaModeEnumOtp:
		return "otp"
	case MfaModeEnumSms:
		return "sms"
	default:
		return "none"
	}
}
