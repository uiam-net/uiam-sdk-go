package uiamsdk

import (
	"errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	goutilsauth "github.com/uiam-net/goutils/auth"
	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

// GenJWTHSToken GenJWTHSToken
func GenJWTHSToken(identityID, realmID, sessionID string, sccret string, duration time.Duration) (string, error) {
	myJWTClaims := uiammodel.AppJWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Audience:  identityID,
			Issuer:    realmID,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	myJWTClaims.SessionID = sessionID

	code, err := goutilsauth.GenerateHS256Token(sccret, &myJWTClaims)

	if err != nil {
		return "", errors.New("Parse Provider Type Error")
	}

	return code, nil
}
