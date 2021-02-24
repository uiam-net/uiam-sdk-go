package uiamsdk

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateHSToken(t *testing.T) {
	token, err := GenJWTHSToken("ae116cae2e2311eb86c0f2189825e567", "808480502e2311ebbfc8f2189825e567", "W$EvF1HCAzj2^8=3n5M0u_fO", "", time.Duration(time.Hour*168))

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf(token)
}
