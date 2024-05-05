// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/FarquestSocial/farquest-go"
	"github.com/FarquestSocial/farquest-go/internal/testutil"
	"github.com/FarquestSocial/farquest-go/option"
)

func TestQuestCompletionCount(t *testing.T) {
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
	err := client.Quest.Completions.Count(context.TODO(), "string")
	if err != nil {
		var apierr *farquest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
