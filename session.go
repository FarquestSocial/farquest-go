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

// SessionService contains methods and other services that help with interacting
// with the farquest API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSessionService] method instead.
type SessionService struct {
	Options []option.RequestOption
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r *SessionService) {
	r = &SessionService{}
	r.Options = opts
	return
}

// Create a new session with a correlated ID
func (r *SessionService) NewCorrelated(ctx context.Context, params SessionNewCorrelatedParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("session/%s", params.PathCorrelatedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type SessionNewCorrelatedParams struct {
	PathCorrelatedID param.Field[string] `path:"correlatedId,required"`
	BodyCorrelatedID param.Field[string] `json:"correlatedId,required"`
	Authorization    param.Field[string] `header:"authorization,required"`
	Farquestapikey   param.Field[string] `header:"farquestapikey,required"`
}

func (r SessionNewCorrelatedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
