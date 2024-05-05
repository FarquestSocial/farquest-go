// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/farquest-go/internal/apijson"
	"github.com/stainless-sdks/farquest-go/internal/param"
	"github.com/stainless-sdks/farquest-go/internal/requestconfig"
	"github.com/stainless-sdks/farquest-go/option"
)

// QuestService contains methods and other services that help with interacting with
// the farquest API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewQuestService] method instead.
type QuestService struct {
	Options     []option.RequestOption
	Types       *QuestTypeService
	Completions *QuestCompletionService
	Validation  *QuestValidationService
}

// NewQuestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuestService(opts ...option.RequestOption) (r *QuestService) {
	r = &QuestService{}
	r.Options = opts
	r.Types = NewQuestTypeService(opts...)
	r.Completions = NewQuestCompletionService(opts...)
	r.Validation = NewQuestValidationService(opts...)
	return
}

// Create a new quest
func (r *QuestService) New(ctx context.Context, body QuestNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "quest/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get a quest by ID, optionally for a specific user by their ID
func (r *QuestService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("quest/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Get quests paginated, optionally with a filter
func (r *QuestService) List(ctx context.Context, filter QuestListParamsFilter, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("quest/list/%v", filter)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type QuestNewParams struct {
	Description            param.Field[string]      `json:"description,required"`
	EndsAt                 param.Field[time.Time]   `json:"endsAt,required" format:"date-time"`
	Image                  param.Field[string]      `json:"image,required"`
	Name                   param.Field[string]      `json:"name,required"`
	OrganizationID         param.Field[string]      `json:"organizationId,required"`
	QuestTypeID            param.Field[string]      `json:"questTypeId,required"`
	StartsAt               param.Field[time.Time]   `json:"startsAt,required" format:"date-time"`
	ValidationCriteria     param.Field[interface{}] `json:"validationCriteria,required"`
	CustomCallbackMetadata param.Field[interface{}] `json:"customCallbackMetadata"`
	CustomMetadata         param.Field[interface{}] `json:"customMetadata"`
}

func (r QuestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QuestListParamsFilter string

const (
	QuestListParamsFilterActive     QuestListParamsFilter = "ACTIVE"
	QuestListParamsFilterComplete   QuestListParamsFilter = "COMPLETE"
	QuestListParamsFilterAll        QuestListParamsFilter = "ALL"
	QuestListParamsFilterNotStarted QuestListParamsFilter = "NOT_STARTED"
)

func (r QuestListParamsFilter) IsKnown() bool {
	switch r {
	case QuestListParamsFilterActive, QuestListParamsFilterComplete, QuestListParamsFilterAll, QuestListParamsFilterNotStarted:
		return true
	}
	return false
}
