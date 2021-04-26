package uiamsdk

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
