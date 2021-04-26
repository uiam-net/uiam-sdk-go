package uiamsdk

import (
	"time"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// IdentityKycRequest IdentityKycRequest
type IdentityKycRequest struct {
	BaseRequest

	IdentityID       string                  `json:"identity_id"`
	Country          uiammodel.CountryEnum   `json:"country"`
	Gender           string                  `json:"gender" binding:"required,min=1"`
	Name             string                  `json:"name" binding:"required,min=1"`
	IDType           uiammodel.IDTypeEnum    `json:"id_type"`
	IDNo             string                  `json:"id_no"`
	Nation           string                  `json:"nation"`
	Birth            string                  `json:"birth"`
	IDValidDateStart string                  `json:"id_valid_date_start"`
	IDValidDateEnd   string                  `json:"id_valid_date_end"`
	EffectedAt       string                  `json:"effected_at"`
	ExpiredAt        string                  `json:"expired_at"`
	Address          string                  `json:"address"`
	IssuedBy         string                  `json:"issued_by"`
	KycLevel         uiammodel.KycLevelEnum  `json:"kyc_level"`
	KycStatus        uiammodel.KycStatusEnum `json:"kyc_status"`
	KycError         string                  `json:"kyc_error"`
	CertifiedAt      *time.Time              `json:"certified_at" binding:"" time_format:"2006-01-02T15:04:05Z07:00"`
}
