package uiamsdk

import (
	"errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	authutil "github.com/uiam-net/goutils/auth"
	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
)

// GenJWTHSToken GenJWTHSToken
func GenJWTHSToken(realmID, sessionID string, sccret string, duration time.Duration) (string, error) {
	myJWTClaims := uiamresp.JWTPayload{
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    realmID,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	myJWTClaims.SessionID = sessionID

	code, err := authutil.GenerateHS256Token(sccret, &myJWTClaims)

	if err != nil {
		return "", errors.New("parse provider type error")
	}

	return code, nil
}
