// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/farquest-go"
	"github.com/stainless-sdks/farquest-go/internal/testutil"
	"github.com/stainless-sdks/farquest-go/option"
)

func TestQuestNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := farquest.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Quest.New(context.TODO(), farquest.QuestNewParams{
		Description:            farquest.F("string"),
		EndsAt:                 farquest.F(time.Now()),
		Image:                  farquest.F("string"),
		Name:                   farquest.F("string"),
		OrganizationID:         farquest.F("string"),
		QuestTypeID:            farquest.F("string"),
		StartsAt:               farquest.F(time.Now()),
		ValidationCriteria:     farquest.F[any](map[string]interface{}{}),
		CustomCallbackMetadata: farquest.F[any](map[string]interface{}{}),
		CustomMetadata:         farquest.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *farquest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuestGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := farquest.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Quest.Get(context.TODO(), "string")
	if err != nil {
		var apierr *farquest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuestListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := farquest.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Quest.List(context.TODO(), farquest.QuestListParamsFilterActive)
	if err != nil {
		var apierr *farquest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
