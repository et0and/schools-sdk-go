// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schools_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/et0and/schools-sdk-go"
	"github.com/et0and/schools-sdk-go/internal/testutil"
	"github.com/et0and/schools-sdk-go/option"
)

func TestRootGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := schools.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Root.Get(context.TODO())
	if err != nil {
		var apierr *schools.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
