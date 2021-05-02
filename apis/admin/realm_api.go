package uiamsdk

import (
	"context"
	"fmt"

	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
	httputil "github.com/uiam-net/uiam-sdk-go/utils/http"
)

// ============ api ============= //
// ============ api ============= //

// GetCurrentRealm GetCurrentRealm
func (ir IdentityRequest) GetCurrentRealm(ctx context.Context) (*uiamresp.Realm, error) {
	var res uiamresp.Realm

	url := fmt.Sprintf("%s/mv1/realm", ir.ServerURL)

	// Request
	if err := httputil.Execute(ir.getRequest(ctx), "GET", url, nil, &res); nil != err {
		return nil, err
	}

	return &res, nil
}
