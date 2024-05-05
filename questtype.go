// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest

import (
	"context"
	"net/http"

	"github.com/FarquestSocial/farquest-go/internal/requestconfig"
	"github.com/FarquestSocial/farquest-go/option"
)

// QuestTypeService contains methods and other services that help with interacting
// with the farquest API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewQuestTypeService] method instead.
type QuestTypeService struct {
	Options []option.RequestOption
}

// NewQuestTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQuestTypeService(opts ...option.RequestOption) (r *QuestTypeService) {
	r = &QuestTypeService{}
	r.Options = opts
	return
}

// Get all quest types
func (r *QuestTypeService) List(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "quest/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
