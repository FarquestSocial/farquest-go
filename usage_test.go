// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package farquest_test

import (
	"context"
	"os"
	"testing"

	"github.com/FarquestSocial/farquest-go"
	"github.com/FarquestSocial/farquest-go/internal/testutil"
	"github.com/FarquestSocial/farquest-go/option"
)

func TestUsage(t *testing.T) {
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
	_, err := client.Organizations.New(context.TODO(), farquest.OrganizationNewParams{
		AuthRedirectURL: farquest.F("string"),
		CallbackURL:     farquest.F("string"),
		Name:            farquest.F("string"),
	})
	if err != nil {
		t.Error(err)
	}
}
