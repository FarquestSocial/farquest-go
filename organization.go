// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/farquest-go/internal/apijson"
	"github.com/stainless-sdks/farquest-go/internal/param"
	"github.com/stainless-sdks/farquest-go/internal/requestconfig"
	"github.com/stainless-sdks/farquest-go/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the farquest API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationService] method
// instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	return
}

// Create a new organization
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "organizations/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type OrganizationNewParams struct {
	AuthRedirectURL param.Field[string] `json:"authRedirectUrl,required"`
	CallbackURL     param.Field[string] `json:"callbackUrl,required"`
	Name            param.Field[string] `json:"name,required"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
