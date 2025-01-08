package tripleseat

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripleseat/utils"
)

func (c *Client) NewGetAllRefundsRequest() GetAllRefundsRequest {
	return GetAllRefundsRequest{
		client:      c,
		queryParams: c.NewGetAllRefundsQueryParams(),
		pathParams:  c.NewGetAllRefundsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetAllRefundsRequestBody(),
	}
}

type GetAllRefundsRequest struct {
	client      *Client
	queryParams *GetAllRefundsQueryParams
	pathParams  *GetAllRefundsPathParams
	method      string
	headers     http.Header
	requestBody GetAllRefundsRequestBody
}

func (c *Client) NewGetAllRefundsQueryParams() *GetAllRefundsQueryParams {
	return &GetAllRefundsQueryParams{}
}

type GetAllRefundsQueryParams struct {
	StartDate             Date         `schema:"start_date,omitempty"`
	EndDate               Date         `schema:"end_date,omitempty"`
	IncludeCategoryTotals BoolToNumber `schema:"include_category_totals,omitempty"`
}

func (p GetAllRefundsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(BoolToNumber(false), utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAllRefundsRequest) QueryParams() *GetAllRefundsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetAllRefundsPathParams() *GetAllRefundsPathParams {
	return &GetAllRefundsPathParams{}
}

type GetAllRefundsPathParams struct {
}

func (p *GetAllRefundsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAllRefundsRequest) PathParams() *GetAllRefundsPathParams {
	return r.pathParams
}

func (r *GetAllRefundsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetAllRefundsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAllRefundsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetAllRefundsRequestBody() GetAllRefundsRequestBody {
	return GetAllRefundsRequestBody{}
}

type GetAllRefundsRequestBody struct {
}

func (r *GetAllRefundsRequest) RequestBody() *GetAllRefundsRequestBody {
	return nil
}

func (r *GetAllRefundsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetAllRefundsRequest) SetRequestBody(body GetAllRefundsRequestBody) {
	r.requestBody = body
}

func (r *GetAllRefundsRequest) NewResponseBody() *GetAllRefundsResponseBody {
	return &GetAllRefundsResponseBody{}
}

type GetAllRefundsResponseBody RefundsAllResp

func (r *GetAllRefundsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("v1/payments/refund_report.json", r.PathParams())
	return &u
}

func (r *GetAllRefundsRequest) Do() (GetAllRefundsResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

// func (r *GetAllRefundsRequest) All() ([]EventItemModel, error) {
// 	AllRefunds := []EventItemModel{}
// 	for {
// 		resp, err := r.Do()
// 		if err != nil {
// 			return AllRefunds, err
// 		}

// 		// Break out of loop when no AllRefunds are found
// 		if len(resp.AllRefunds) == 0 {
// 			break
// 		}

// 		// Add AllRefunds to list
// 		AllRefunds = append(AllRefunds, resp.AllRefunds...)

// 		if len(AllRefunds) == resp.Count {
// 			break
// 		}

// 		// Increment page number
// 		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
// 	}

// 	return AllRefunds, nil
// }
