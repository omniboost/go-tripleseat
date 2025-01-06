package tripleseat

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-tripleseat/utils"
)

func (c *Client) NewGetEventByIdRequest() GetEventByIdRequest {
	return GetEventByIdRequest{
		client:      c,
		queryParams: c.NewGetEventByIdQueryParams(),
		pathParams:  c.NewGetEventByIdPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetEventByIdRequestBody(),
	}
}

type GetEventByIdRequest struct {
	client      *Client
	queryParams *GetEventByIdQueryParams
	pathParams  *GetEventByIdPathParams
	method      string
	headers     http.Header
	requestBody GetEventByIdRequestBody
}

func (c *Client) NewGetEventByIdQueryParams() *GetEventByIdQueryParams {
	return &GetEventByIdQueryParams{}
}

type GetEventByIdQueryParams struct {
	ShowFinancials bool `schema:"show_financials,omitempty"`
}

func (p GetEventByIdQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetEventByIdRequest) QueryParams() *GetEventByIdQueryParams {
	return r.queryParams
}

func (c *Client) NewGetEventByIdPathParams() *GetEventByIdPathParams {
	return &GetEventByIdPathParams{}
}

type GetEventByIdPathParams struct {
	EventID int `schema:"event_id,omitempty"`
}

func (p *GetEventByIdPathParams) Params() map[string]string {
	return map[string]string{
		"event_id": strconv.Itoa(p.EventID),
	}
}

func (r *GetEventByIdRequest) PathParams() *GetEventByIdPathParams {
	return r.pathParams
}

func (r *GetEventByIdRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetEventByIdRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetEventByIdRequest) Method() string {
	return r.method
}

func (s *Client) NewGetEventByIdRequestBody() GetEventByIdRequestBody {
	return GetEventByIdRequestBody{}
}

type GetEventByIdRequestBody struct {
}

func (r *GetEventByIdRequest) RequestBody() *GetEventByIdRequestBody {
	return nil
}

func (r *GetEventByIdRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetEventByIdRequest) SetRequestBody(body GetEventByIdRequestBody) {
	r.requestBody = body
}

func (r *GetEventByIdRequest) NewResponseBody() *GetEventByIdResponseBody {
	return &GetEventByIdResponseBody{}
}

type GetEventByIdResponseBody EventResp

func (r *GetEventByIdRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("v1/events/{{.event_id}}", r.PathParams())
	return &u
}

func (r *GetEventByIdRequest) Do() (GetEventByIdResponseBody, error) {
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

// func (r *GetEventByIdRequest) All() ([]EventItemModel, error) {
// 	EventById := []EventItemModel{}
// 	for {
// 		resp, err := r.Do()
// 		if err != nil {
// 			return EventById, err
// 		}

// 		// Break out of loop when no EventById are found
// 		if len(resp.EventById) == 0 {
// 			break
// 		}

// 		// Add EventById to list
// 		EventById = append(EventById, resp.EventById...)

// 		if len(EventById) == resp.Count {
// 			break
// 		}

// 		// Increment page number
// 		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
// 	}

// 	return EventById, nil
// }
