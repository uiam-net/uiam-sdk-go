package uiamsdk

// =============== IdentityTypeEnum =============== //

// IdentityTypeEnum 枚举
type IdentityTypeEnum string

const (
	IdentityTypeEnumStd IdentityTypeEnum = "std"
	IdentityTypeEnumSys IdentityTypeEnum = "sys"
	IdentityTypeEnumVir IdentityTypeEnum = "vir"
)

func (e IdentityTypeEnum) String() string {
	switch e {
	case IdentityTypeEnumStd:
		return "std"
	case IdentityTypeEnumSys:
		return "sys"
	default:
		return "unkonwn"
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
		return "unkonwn"
	}
}
