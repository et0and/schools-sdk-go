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

func TestSchoolGet(t *testing.T) {
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
	_, err := client.Schools.Get(context.TODO(), "schoolId")
	if err != nil {
		var apierr *schools.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSchoolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Schools.List(context.TODO(), schools.SchoolListParams{
		Authority: schools.String("authority"),
		City:      schools.String("city"),
		Limit:     schools.Int(1),
		Name:      schools.String("name"),
		OrgType:   schools.String("org_type"),
		Page:      schools.Int(1),
		Status:    schools.String("status"),
		Suburb:    schools.String("suburb"),
	})
	if err != nil {
		var apierr *schools.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSchoolByAuthorityWithOptionalParams(t *testing.T) {
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
	err := client.Schools.ByAuthority(
		context.TODO(),
		"authority",
		schools.SchoolByAuthorityParams{
			Limit: schools.Int(1),
			Page:  schools.Int(1),
		},
	)
	if err != nil {
		var apierr *schools.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSchoolByCityWithOptionalParams(t *testing.T) {
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
	err := client.Schools.ByCity(
		context.TODO(),
		"city",
		schools.SchoolByCityParams{
			Limit: schools.Int(1),
			Page:  schools.Int(1),
		},
	)
	if err != nil {
		var apierr *schools.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSchoolByStatusWithOptionalParams(t *testing.T) {
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
	err := client.Schools.ByStatus(
		context.TODO(),
		"status",
		schools.SchoolByStatusParams{
			Limit: schools.Int(1),
			Page:  schools.Int(1),
		},
	)
	if err != nil {
		var apierr *schools.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSchoolBySuburbWithOptionalParams(t *testing.T) {
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
	err := client.Schools.BySuburb(
		context.TODO(),
		"suburb",
		schools.SchoolBySuburbParams{
			Limit: schools.Int(1),
			Page:  schools.Int(1),
		},
	)
	if err != nil {
		var apierr *schools.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSchoolSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Schools.Search(context.TODO(), schools.SchoolSearchParams{
		Q:     "x",
		Limit: schools.Int(1),
		Page:  schools.Int(1),
	})
	if err != nil {
		var apierr *schools.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
