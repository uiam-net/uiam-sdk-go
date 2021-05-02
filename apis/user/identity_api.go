package uiamsdk

import (
	"context"
	"fmt"
	"strings"

	uiamresp "github.com/uiam-net/uiam-sdk-go/responses"
	httputil "github.com/uiam-net/uiam-sdk-go/utils/http"
)

// ============ api ============= //
// ============ api ============= //

// GetCurrentIdentity GetUser
func (r IdentityRequest) GetCurrentIdentity(ctx context.Context, profile bool) (*uiamresp.Identity, error) {
	var res uiamresp.Identity

	var expand = make([]string, 0)
	if profile {
		expand = append(expand, "profile")
	}

	url := fmt.Sprintf("%s/uv1/identity?expand=%s", r.ServerURL, strings.Join(expand, ","))

	// Request
	if err := httputil.Execute(r.getRequest(ctx), "GET", url, nil, &res); nil != err {
		return nil, err
	}

	return &res, nil
}
