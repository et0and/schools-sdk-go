// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schools

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/et0and/schools-sdk-go/internal/apijson"
	"github.com/et0and/schools-sdk-go/internal/requestconfig"
	"github.com/et0and/schools-sdk-go/option"
	"github.com/et0and/schools-sdk-go/packages/respjson"
)

// SyncService contains methods and other services that help with interacting with
// the schools API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSyncService] method instead.
type SyncService struct {
	Options []option.RequestOption
}

// NewSyncService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSyncService(opts ...option.RequestOption) (r SyncService) {
	r = SyncService{}
	r.Options = opts
	return
}

// Get sync status
func (r *SyncService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *SyncGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sync/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Trigger manual data sync
func (r *SyncService) Trigger(ctx context.Context, opts ...option.RequestOption) (res *SyncTriggerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sync"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SyncGetStatusResponse struct {
	IsStale     bool      `json:"isStale"`
	LastSync    time.Time `json:"lastSync,nullable" format:"date-time"`
	RecordCount int64     `json:"recordCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsStale     respjson.Field
		LastSync    respjson.Field
		RecordCount respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyncGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *SyncGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SyncTriggerResponse struct {
	Error       string    `json:"error"`
	LastSync    time.Time `json:"lastSync" format:"date-time"`
	RecordCount int64     `json:"recordCount"`
	Success     bool      `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		LastSync    respjson.Field
		RecordCount respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyncTriggerResponse) RawJSON() string { return r.JSON.raw }
func (r *SyncTriggerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
