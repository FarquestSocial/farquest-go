// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/farquest-go"
	"github.com/stainless-sdks/farquest-go/internal/testutil"
	"github.com/stainless-sdks/farquest-go/option"
)

func TestQuestTypeList(t *testing.T) {
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
	err := client.Quest.Types.List(context.TODO())
	if err != nil {
		var apierr *farquest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
