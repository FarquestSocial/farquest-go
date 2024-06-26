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

func TestSessionNewCorrelated(t *testing.T) {
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
	err := client.Session.NewCorrelated(context.TODO(), farquest.SessionNewCorrelatedParams{
		PathCorrelatedID: farquest.F("string"),
		BodyCorrelatedID: farquest.F("string"),
		Authorization:    farquest.F("string"),
		Farquestapikey:   farquest.F("string"),
	})
	if err != nil {
		var apierr *farquest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
