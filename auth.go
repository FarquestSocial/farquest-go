// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FarquestSocial/farquest-go/internal/apijson"
	"github.com/FarquestSocial/farquest-go/internal/param"
	"github.com/FarquestSocial/farquest-go/internal/requestconfig"
	"github.com/FarquestSocial/farquest-go/option"
)

// AuthService contains methods and other services that help with interacting with
// the farquest API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r *AuthService) {
	r = &AuthService{}
	r.Options = opts
	return
}

// Callback to session with Privy Auth, returns URL to redirect back to your app
func (r *AuthService) Callback(ctx context.Context, body AuthCallbackParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "auth/callback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get the redirect URL to start a session
func (r *AuthService) GetState(ctx context.Context, state string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("auth/%s", state)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type AuthCallbackParams struct {
	State param.Field[string] `json:"state,required"`
}

func (r AuthCallbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
