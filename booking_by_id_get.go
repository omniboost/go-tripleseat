package tripleseat

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-tripleseat/utils"
)

func (c *Client) NewGetBookingByIdRequest() GetBookingByIdRequest {
	return GetBookingByIdRequest{
		client:      c,
		queryParams: c.NewGetBookingByIdQueryParams(),
		pathParams:  c.NewGetBookingByIdPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetBookingByIdRequestBody(),
	}
}

type GetBookingByIdRequest struct {
	client      *Client
	queryParams *GetBookingByIdQueryParams
	pathParams  *GetBookingByIdPathParams
	method      string
	headers     http.Header
	requestBody GetBookingByIdRequestBody
}

func (c *Client) NewGetBookingByIdQueryParams() *GetBookingByIdQueryParams {
	return &GetBookingByIdQueryParams{}
}

type GetBookingByIdQueryParams struct {
	ShowFinancials bool `schema:"show_financials,omitempty"`
}

func (p GetBookingByIdQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetBookingByIdRequest) QueryParams() *GetBookingByIdQueryParams {
	return r.queryParams
}

func (c *Client) NewGetBookingByIdPathParams() *GetBookingByIdPathParams {
	return &GetBookingByIdPathParams{}
}

type GetBookingByIdPathParams struct {
	BookingID int `schema:"booking_id,omitempty"`
}

func (p *GetBookingByIdPathParams) Params() map[string]string {
	return map[string]string{
		"booking_id": strconv.Itoa(p.BookingID),
	}
}

func (r *GetBookingByIdRequest) PathParams() *GetBookingByIdPathParams {
	return r.pathParams
}

func (r *GetBookingByIdRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetBookingByIdRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetBookingByIdRequest) Method() string {
	return r.method
}

func (s *Client) NewGetBookingByIdRequestBody() GetBookingByIdRequestBody {
	return GetBookingByIdRequestBody{}
}

type GetBookingByIdRequestBody struct {
}

func (r *GetBookingByIdRequest) RequestBody() *GetBookingByIdRequestBody {
	return nil
}

func (r *GetBookingByIdRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetBookingByIdRequest) SetRequestBody(body GetBookingByIdRequestBody) {
	r.requestBody = body
}

func (r *GetBookingByIdRequest) NewResponseBody() *GetBookingByIdResponseBody {
	return &GetBookingByIdResponseBody{}
}

type GetBookingByIdResponseBody BookingResp

func (r *GetBookingByIdRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("v1/bookings/{{.booking_id}}", r.PathParams())
	return &u
}

func (r *GetBookingByIdRequest) Do() (GetBookingByIdResponseBody, error) {
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

// func (r *GetBookingByIdRequest) All() ([]BookingItemModel, error) {
// 	BookingById := []BookingItemModel{}
// 	for {
// 		resp, err := r.Do()
// 		if err != nil {
// 			return BookingById, err
// 		}

// 		// Break out of loop when no BookingById are found
// 		if len(resp.BookingById) == 0 {
// 			break
// 		}

// 		// Add BookingById to list
// 		BookingById = append(BookingById, resp.BookingById...)

// 		if len(BookingById) == resp.Count {
// 			break
// 		}

// 		// Increment page number
// 		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
// 	}

// 	return BookingById, nil
// }
