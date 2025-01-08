package tripleseat

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripleseat/utils"
)

func (c *Client) NewGetAllPaymentsRequest() GetAllPaymentsRequest {
	return GetAllPaymentsRequest{
		client:      c,
		queryParams: c.NewGetAllPaymentsQueryParams(),
		pathParams:  c.NewGetAllPaymentsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetAllPaymentsRequestBody(),
	}
}

type GetAllPaymentsRequest struct {
	client      *Client
	queryParams *GetAllPaymentsQueryParams
	pathParams  *GetAllPaymentsPathParams
	method      string
	headers     http.Header
	requestBody GetAllPaymentsRequestBody
}

func (c *Client) NewGetAllPaymentsQueryParams() *GetAllPaymentsQueryParams {
	return &GetAllPaymentsQueryParams{}
}

type GetAllPaymentsQueryParams struct {
	StartDate             Date         `schema:"start_date,omitempty"`
	EndDate               Date         `schema:"end_date,omitempty"`
	IncludeCategoryTotals BoolToNumber `schema:"include_category_totals,omitempty"`
}

func (p GetAllPaymentsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetAllPaymentsRequest) QueryParams() *GetAllPaymentsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetAllPaymentsPathParams() *GetAllPaymentsPathParams {
	return &GetAllPaymentsPathParams{}
}

type GetAllPaymentsPathParams struct {
}

func (p *GetAllPaymentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAllPaymentsRequest) PathParams() *GetAllPaymentsPathParams {
	return r.pathParams
}

func (r *GetAllPaymentsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetAllPaymentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAllPaymentsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetAllPaymentsRequestBody() GetAllPaymentsRequestBody {
	return GetAllPaymentsRequestBody{}
}

type GetAllPaymentsRequestBody struct {
}

func (r *GetAllPaymentsRequest) RequestBody() *GetAllPaymentsRequestBody {
	return nil
}

func (r *GetAllPaymentsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetAllPaymentsRequest) SetRequestBody(body GetAllPaymentsRequestBody) {
	r.requestBody = body
}

func (r *GetAllPaymentsRequest) NewResponseBody() *GetAllPaymentsResponseBody {
	return &GetAllPaymentsResponseBody{}
}

type GetAllPaymentsResponseBody PaymentsAllResp

func (r *GetAllPaymentsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("v1/payments/transaction_report.json", r.PathParams())
	return &u
}

func (r *GetAllPaymentsRequest) Do() (GetAllPaymentsResponseBody, error) {
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

// func (r *GetAllPaymentsRequest) All() ([]EventItemModel, error) {
// 	AllPayments := []EventItemModel{}
// 	for {
// 		resp, err := r.Do()
// 		if err != nil {
// 			return AllPayments, err
// 		}

// 		// Break out of loop when no AllPayments are found
// 		if len(resp.AllPayments) == 0 {
// 			break
// 		}

// 		// Add AllPayments to list
// 		AllPayments = append(AllPayments, resp.AllPayments...)

// 		if len(AllPayments) == resp.Count {
// 			break
// 		}

// 		// Increment page number
// 		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
// 	}

// 	return AllPayments, nil
// }
