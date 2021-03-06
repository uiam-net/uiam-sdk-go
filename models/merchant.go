package uiamsdk

// =============== IdentityTypeEnum =============== //

type MerchantTypeEnum string

const (
	MerchantTypeEnumPersonal MerchantTypeEnum = "personal" // 正常注册用户
	MerchantTypeEnumCompany  MerchantTypeEnum = "copmpany" // 系统调用使用
)

func (e MerchantTypeEnum) String() string {
	switch e {
	case MerchantTypeEnumPersonal:
		return "personal"
	case MerchantTypeEnumCompany:
		return "copmpany"
	default:
		return ""
	}
}

// MerchantStatusEnum 枚举
type MerchantStatusEnum string

const (
	// MerchantStatusEnumNormal 正常
	MerchantStatusEnumNormal MerchantStatusEnum = "normal"
	// MerchantStatusEnumDisabled 下架
	MerchantStatusEnumDisabled MerchantStatusEnum = "disabled"
)

func (e MerchantStatusEnum) String() string {
	switch e {
	case MerchantStatusEnumNormal:
		return "normal"
	case MerchantStatusEnumDisabled:
		return "disabled"
	default:
		return ""
	}
}
