package uiamsdk

// ================== EUUM ================== //

// RealmStatusEnum 枚举
type RealmAuthzModelEnum string

const (
	RealmAuthzModelEnumRealm      RealmAuthzModelEnum = "realm"
	RealmAuthzModelEnumUser       RealmAuthzModelEnum = "user"
	RealmAuthzModelEnumUserDevice RealmAuthzModelEnum = "device"
)

func (e RealmAuthzModelEnum) String() string {
	switch e {
	case RealmAuthzModelEnumRealm:
		return "realm"
	case RealmAuthzModelEnumUser:
		return "user"
	case RealmAuthzModelEnumUserDevice:
		return "device"
	default:
		return ""
	}
}

// RealmStatusEnum 枚举
type RealmStatusEnum string

const (
	// RealmStatusEnumNormal 正常
	RealmStatusEnumNormal RealmStatusEnum = "normal"
	// RealmStatusEnumDisabled 下架
	RealmStatusEnumDisabled RealmStatusEnum = "disabled"
)

func (e RealmStatusEnum) String() string {
	switch e {
	case RealmStatusEnumNormal:
		return "normal"
	case RealmStatusEnumDisabled:
		return "disabled"
	default:
		return ""
	}
}

// RealmCaptchaModeEnum 枚举
type RealmCaptchaModeEnum string

const (
	RealmCaptchaModeEnumNone RealmCaptchaModeEnum = "none" // 禁用
	RealmCaptchaModeEnumPre  RealmCaptchaModeEnum = "pre"  // 前置，在提交登录请求时，需要带上 captcha 数据
	RealmCaptchaModeEnumPost RealmCaptchaModeEnum = "post" // 后置，提交完登录请求后，再根据情况判断是否需要识别
)

func (e RealmCaptchaModeEnum) String() string {
	switch e {
	case RealmCaptchaModeEnumNone:
		return "none"
	case RealmCaptchaModeEnumPre:
		return "pre"
	case RealmCaptchaModeEnumPost:
		return "post"
	default:
		return ""
	}
}

// RealmCaptchaProviderEnum 枚举
type RealmCaptchaProviderEnum string

const (
	RealmCaptchaProviderEnumSMS     RealmCaptchaProviderEnum = "sms"
	RealmCaptchaProviderEnumUiam    RealmCaptchaProviderEnum = "uiam"
	RealmCaptchaProviderEnumAliyun  RealmCaptchaProviderEnum = "aliyun"
	RealmCaptchaProviderEnumGeetest RealmCaptchaProviderEnum = "geetest"
)

func (e RealmCaptchaProviderEnum) String() string {
	switch e {
	case RealmCaptchaProviderEnumSMS:
		return "sms"
	case RealmCaptchaProviderEnumUiam:
		return "uiam"
	case RealmCaptchaProviderEnumAliyun:
		return "aliyun"
	case RealmCaptchaProviderEnumGeetest:
		return "geetest"
	default:
		return "none"
	}
}

// RealmCaptchaTypeEnum 枚举
type RealmCaptchaTypeEnum string

const (
	RealmCaptchaTypeEnumDigtal          RealmCaptchaTypeEnum = "digtal"
	RealmCaptchaTypeEnumString          RealmCaptchaTypeEnum = "string"
	RealmCaptchaTypeEnumMath            RealmCaptchaTypeEnum = "math"
	RealmCaptchaTypeEnumSlide           RealmCaptchaTypeEnum = "slide"
	RealmCaptchaTypeEnumBlockPuzzle     RealmCaptchaTypeEnum = "block_puzzle"
	RealmCaptchaTypeEnumClickWord       RealmCaptchaTypeEnum = "click_word"
	RealmCaptchaTypeEnumGoogleRecaptcha RealmCaptchaTypeEnum = "recaptcha"
)

func (e RealmCaptchaTypeEnum) String() string {
	switch e {
	case RealmCaptchaTypeEnumDigtal:
		return "digtal"
	case RealmCaptchaTypeEnumString:
		return "string"
	case RealmCaptchaTypeEnumMath:
		return "math"
	case RealmCaptchaTypeEnumSlide:
		return "slide"
	case RealmCaptchaTypeEnumBlockPuzzle:
		return "block_puzzle"
	case RealmCaptchaTypeEnumClickWord:
		return "click_word"
	default:
		return ""
	}
}
