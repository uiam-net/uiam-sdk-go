package uiamsdk

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
