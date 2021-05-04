package uiamsdk

import (
	authutil "github.com/uiam-net/goutils/auth"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// ==== Verify Operations ===== //
// ==== Verify Operations ===== //

// VerifyClaimsHS VerifyClaimsHS
func VerifyClaimsHS(token, secret string, payload jwtgo.Claims) error {
	return authutil.VerifyClaimsHS(token, secret, payload)
}

// VerifyClaimsRS VerifyClaimsRS
func VerifyClaimsRS(tokenStr, pubkey string, payload jwtgo.Claims) error {
	return authutil.VerifyClaimsRS(tokenStr, pubkey, payload)
}

// ==== Parse Operations ===== //
// ==== Parse Operations ===== //

// ParseTokenUnverified ParseTokenUnverified
func ParseJWTUnverified(tokenStr string, payload jwtgo.Claims) (*jwtgo.Token, error) {
	return authutil.ParseTokenUnverified(tokenStr, payload)
}

// ParseClaimsUnverified ParseClaimsUnverified
func ParseClaimsUnverified(tokenStr string, payload jwtgo.Claims) (*jwtgo.Token, jwtgo.Claims, error) {
	return authutil.ParseClaimsUnverified(tokenStr, payload)
}

// ==== Generate Operations ===== //
// ==== Generate Operations ===== //

// GenerateHS256Token GenerateHS256Token
func GenerateHS256Token(secret string, payload jwtgo.Claims) (string, error) {
	return authutil.GenerateHS256Token(secret, payload)
}

// GenerateRSA256Token GenerateRSAToken
func GenerateRSA256Token(prikey string, payload jwtgo.Claims) (string, error) {
	return authutil.GenerateRSA256Token(prikey, payload)
}

// GenerateRSA512Token GenerateRSAToken
func GenerateRSA512Token(appID, prikey string, payload jwtgo.Claims) (string, error) {
	return authutil.GenerateRSA512Token(appID, prikey, payload)
}
