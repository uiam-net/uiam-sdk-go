package uiamsdk

// =============== KycStatusEnum =============== //

// KycStatusEnum 枚举
type KycStatusEnum string

const (
	KycStatusEnumProcessing KycStatusEnum = "processing"
	KycStatusEnumRejected   KycStatusEnum = "rejected"
	KycStatusEnumApproved   KycStatusEnum = "approved"
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
		return ""
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
		return ""
	}
}
