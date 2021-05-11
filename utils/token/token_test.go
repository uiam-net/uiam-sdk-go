package uiamsdk

import (
	"fmt"
	"testing"
	"time"

	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
)

func TestGenerateHSToken(t *testing.T) {
	// Build JWT
	myJWTClaims := uiamresp.JWTPayload{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(time.Hour * 168)).Unix(),
	}

	// 系统数据
	myJWTClaims.Issuer = "realmid"
	myJWTClaims.Subject = "userid" // 用系统的值，不然会有问题

	// 客户自定义数据
	myJWTClaims.Id = "jti" // JTI 由用户端生成才有意义

	token, err := GenerateHS256Token("ae116cae2e2311eb86c0f2189825e567", myJWTClaims)

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf(token)
}
