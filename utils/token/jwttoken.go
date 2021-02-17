package uiamsdk

import (
	"errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	goutilsauth "github.com/uiam-net/goutils/auth"
	uiammodels "github.com/uiam-net/uiam-sdk-go/models"
)

// GenJWTHSToken GenJWTHSToken
func GenJWTHSToken(identityID, realmID string, sccret string, duration time.Duration) (string, error) {
	myJWTClaims := uiammodels.AppJWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Audience:  identityID,
			Issuer:    realmID,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	code, err := goutilsauth.GenerateHS256Token(sccret, &myJWTClaims)

	if err != nil {
		return "", errors.New("Parse Provider Type Error")
	}

	return code, nil
}
