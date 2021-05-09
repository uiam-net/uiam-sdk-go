package uiamsdk

import (
	"encoding/json"
	"fmt"
	"testing"

	uiammodel "github.com/uiam-net/uiam-sdk-go/models"
)

func TestGenerateHSToken(t *testing.T) {
	// Build JWT
	payload := CaptchaPayloadUiam{
		Type:   uiammodel.RealmCaptchaTypeEnumDigtal,
		Length: 4,
		Width:  320,
		Height: 40,
	}

	aaa, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%s", string(aaa))
}
