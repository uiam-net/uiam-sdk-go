package uiamsdk

import (
	"fmt"
	"time"

	aesutils "github.com/uiam-net/goutils/crypto/aes"
	hashutil "github.com/uiam-net/goutils/crypto/hash"
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
	uiamerr "github.com/uiam-net/uiam-sdk-go/utils/error"
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

// NewProfileRequest NewProfile
func NewProfileRequest(request *IdentityKycRequest, aesScecret string) (*uiammodel.Profile, uiamerr.Error) {
	var idNOEncrypted, idDigest, birthEncrypted, issuedByEncrypted, addressEncrypted, effectedAt, expiredAt string
	var err error

	if request.IDNo != "" {
		idNOEncrypted, err = aesutils.Encrypt([]byte(request.IDNo), []byte(aesScecret), []byte(request.IdentityID))

		if err != nil {
			return nil, uiamerr.NewAppError(err.Error())
		}

		idHashContent := fmt.Sprintf("CN-IDCARD-%s", request.IDNo)
		idDigest = hashutil.SHA1Hex(idHashContent)
	}

	if request.Birth != "" {
		birthEncrypted, err = aesutils.Encrypt([]byte(request.Birth), []byte(aesScecret), []byte(request.IdentityID))

		if err != nil {
			return nil, uiamerr.NewAppError(err.Error())
		}
	}

	if request.IssuedBy != "" {
		issuedByEncrypted, err = aesutils.Encrypt([]byte(request.IssuedBy), []byte(aesScecret), []byte(request.IdentityID))

		if err != nil {
			return nil, uiamerr.NewAppError(err.Error())
		}
	}

	if request.Address != "" {
		addressEncrypted, err = aesutils.Encrypt([]byte(request.Address), []byte(aesScecret), []byte(request.IdentityID))

		if err != nil {
			return nil, uiamerr.NewAppError(err.Error())
		}
	}

	if request.EffectedAt != "" {
		effectedAt = request.EffectedAt
	} else {
		effectedAt = request.IDValidDateStart
	}

	if request.ExpiredAt != "" {
		expiredAt = request.ExpiredAt
	} else {
		expiredAt = request.IDValidDateEnd
	}

	genderEnum := uiammodel.NewGenderEnum(request.Gender)

	if request.CertifiedAt == nil {
		timeNow := time.Now()
		request.CertifiedAt = &timeNow
	}

	newProfile := &uiammodel.Profile{
		IdentityID:        request.IdentityID,
		Name:              request.Name,
		Gender:            genderEnum,
		BirthEncrypted:    birthEncrypted,
		Country:           request.Country,
		IDType:            request.IDType,
		IDNoEncrypted:     idNOEncrypted,
		IDDigest:          idDigest,
		Nation:            request.Nation,
		EffectedAt:        effectedAt,
		ExpiredAt:         expiredAt,
		IssuedByEncrypted: issuedByEncrypted,
		AddressEncrypted:  addressEncrypted,
		KycLevel:          request.KycLevel,
		KycStatus:         request.KycStatus,
		KycError:          request.KycError,
		CertifiedAt:       request.CertifiedAt,
	}

	return newProfile, nil
}
