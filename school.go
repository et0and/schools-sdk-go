// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schools

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/et0and/schools-sdk-go/internal/apijson"
	"github.com/et0and/schools-sdk-go/internal/apiquery"
	"github.com/et0and/schools-sdk-go/internal/requestconfig"
	"github.com/et0and/schools-sdk-go/option"
	"github.com/et0and/schools-sdk-go/packages/param"
	"github.com/et0and/schools-sdk-go/packages/respjson"
)

// School data retrieval operations
//
// SchoolService contains methods and other services that help with interacting
// with the schools API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSchoolService] method instead.
type SchoolService struct {
	Options []option.RequestOption
}

// NewSchoolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSchoolService(opts ...option.RequestOption) (r SchoolService) {
	r = SchoolService{}
	r.Options = opts
	return
}

// Get school by School ID
func (r *SchoolService) Get(ctx context.Context, schoolID string, opts ...option.RequestOption) (res *SchoolGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if schoolID == "" {
		err = errors.New("missing required schoolId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/schools/id/%s", schoolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get all schools with filtering
func (r *SchoolService) List(ctx context.Context, query SchoolListParams, opts ...option.RequestOption) (res *SchoolListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/schools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get schools by authority
func (r *SchoolService) ByAuthority(ctx context.Context, authority string, query SchoolByAuthorityParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if authority == "" {
		err = errors.New("missing required authority parameter")
		return err
	}
	path := fmt.Sprintf("v1/schools/authority/%s", authority)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Get schools by city
func (r *SchoolService) ByCity(ctx context.Context, city string, query SchoolByCityParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if city == "" {
		err = errors.New("missing required city parameter")
		return err
	}
	path := fmt.Sprintf("v1/schools/city/%s", city)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Get schools by status
func (r *SchoolService) ByStatus(ctx context.Context, status string, query SchoolByStatusParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if status == "" {
		err = errors.New("missing required status parameter")
		return err
	}
	path := fmt.Sprintf("v1/schools/status/%s", status)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Get schools by suburb
func (r *SchoolService) BySuburb(ctx context.Context, suburb string, query SchoolBySuburbParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if suburb == "" {
		err = errors.New("missing required suburb parameter")
		return err
	}
	path := fmt.Sprintf("v1/schools/suburb/%s", suburb)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Full-text search schools by name
func (r *SchoolService) Search(ctx context.Context, query SchoolSearchParams, opts ...option.RequestOption) (res *SchoolSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/schools/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SchoolGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchoolGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SchoolGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchoolListResponse struct {
	Data       []any                        `json:"data"`
	Pagination SchoolListResponsePagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchoolListResponse) RawJSON() string { return r.JSON.raw }
func (r *SchoolListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchoolListResponsePagination struct {
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchoolListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SchoolListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchoolSearchResponse struct {
	Data       []any                          `json:"data"`
	Pagination SchoolSearchResponsePagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchoolSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SchoolSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchoolSearchResponsePagination struct {
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchoolSearchResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SchoolSearchResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchoolListParams struct {
	// Filter by education authority
	Authority param.Opt[string] `query:"authority,omitzero" json:"-"`
	// Filter by city (partial match)
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Results per page (default: 20, max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by school name (partial match)
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by organization type
	OrgType param.Opt[string] `query:"org_type,omitzero" json:"-"`
	// Page number (default: 1)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by school status
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter by suburb (partial match)
	Suburb param.Opt[string] `query:"suburb,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SchoolListParams]'s query parameters as `url.Values`.
func (r SchoolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SchoolByAuthorityParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SchoolByAuthorityParams]'s query parameters as
// `url.Values`.
func (r SchoolByAuthorityParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SchoolByCityParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SchoolByCityParams]'s query parameters as `url.Values`.
func (r SchoolByCityParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SchoolByStatusParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SchoolByStatusParams]'s query parameters as `url.Values`.
func (r SchoolByStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SchoolBySuburbParams struct {
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page  param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SchoolBySuburbParams]'s query parameters as `url.Values`.
func (r SchoolBySuburbParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SchoolSearchParams struct {
	// Search query
	Q string `query:"q" api:"required" json:"-"`
	// Results per page (default: 20, max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number (default: 1)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SchoolSearchParams]'s query parameters as `url.Values`.
func (r SchoolSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
