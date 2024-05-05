// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/farquest-go/internal/requestconfig"
	"github.com/stainless-sdks/farquest-go/option"
)

// QuestValidationService contains methods and other services that help with
// interacting with the farquest API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewQuestValidationService] method
// instead.
type QuestValidationService struct {
	Options []option.RequestOption
}

// NewQuestValidationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQuestValidationService(opts ...option.RequestOption) (r *QuestValidationService) {
	r = &QuestValidationService{}
	r.Options = opts
	return
}

// Get the validation criteria for a quest
func (r *QuestValidationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("quest/validation/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
