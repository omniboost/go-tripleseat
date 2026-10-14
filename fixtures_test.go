package tripleseat_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	tripleseat "github.com/omniboost/go-tripleseat"
)

// newTestClient points a Client at a local httptest server instead of the live
// Tripleseat API. It uses the default HTTP client (no OAuth) and leaves debug
// disabled, so these tests run offline and make no authenticated calls.
func newTestClient(t *testing.T, srv *httptest.Server) *tripleseat.Client {
	t.Helper()
	u, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatalf("parse test server URL: %v", err)
	}
	c := tripleseat.NewClient(nil)
	c.SetBaseURL(*u)
	return c
}

// jsonHandler returns a handler that records the request path and raw query,
// then responds with the given JSON body and the application/json content type
// the client's CheckResponse requires.
func jsonHandler(gotPath, gotQuery *string, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		*gotPath = r.URL.Path
		*gotQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	}
}

func TestGetEventByIdRequestOffline(t *testing.T) {
	var gotPath, gotQuery string
	srv := httptest.NewServer(jsonHandler(&gotPath, &gotQuery,
		`{"event":{"id":45103683,"name":"Test Event","status":"definite"}}`))
	defer srv.Close()

	c := newTestClient(t, srv)
	req := c.NewGetEventByIdRequest()
	req.PathParams().EventID = 45103683
	req.QueryParams().ShowFinancials = true

	resp, err := req.Do()
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if resp.Event.ID != 45103683 {
		t.Errorf("Event.ID = %d, want 45103683", resp.Event.ID)
	}
	if resp.Event.Name != "Test Event" {
		t.Errorf("Event.Name = %q, want %q", resp.Event.Name, "Test Event")
	}
	if want := "/v1/events/45103683"; gotPath != want {
		t.Errorf("request path = %q, want %q", gotPath, want)
	}
	if !strings.Contains(gotQuery, "show_financial=true") {
		t.Errorf("query = %q, want it to contain show_financial=true", gotQuery)
	}
	if strings.Contains(gotQuery, "show_financials") {
		t.Errorf("query = %q, should not use the old show_financials tag", gotQuery)
	}
}

func TestGetBookingByIdRequestOffline(t *testing.T) {
	var gotPath, gotQuery string
	srv := httptest.NewServer(jsonHandler(&gotPath, &gotQuery,
		`{"booking":{"id":42946736,"name":"Test Booking","status":"definite"}}`))
	defer srv.Close()

	c := newTestClient(t, srv)
	req := c.NewGetBookingByIdRequest()
	req.PathParams().BookingID = 42946736
	req.QueryParams().ShowFinancials = true

	resp, err := req.Do()
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if resp.Booking.ID != 42946736 {
		t.Errorf("Booking.ID = %d, want 42946736", resp.Booking.ID)
	}
	if resp.Booking.Name != "Test Booking" {
		t.Errorf("Booking.Name = %q, want %q", resp.Booking.Name, "Test Booking")
	}
	if want := "/v1/bookings/42946736"; gotPath != want {
		t.Errorf("request path = %q, want %q", gotPath, want)
	}
	if !strings.Contains(gotQuery, "show_financial=true") {
		t.Errorf("query = %q, want it to contain show_financial=true", gotQuery)
	}
	if strings.Contains(gotQuery, "show_financials") {
		t.Errorf("query = %q, should not use the old show_financials tag", gotQuery)
	}
}

func TestGetAllPaymentsRequestOffline(t *testing.T) {
	var gotPath, gotQuery string
	srv := httptest.NewServer(jsonHandler(&gotPath, &gotQuery,
		`{"success":true,"data":[{"rowid":"pmt-1","payment_method":"credit_card","amount":"100.0"}]}`))
	defer srv.Close()

	c := newTestClient(t, srv)
	req := c.NewGetAllPaymentsRequest()
	req.QueryParams().StartDate = tripleseat.Date{time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().EndDate = tripleseat.Date{time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().IncludeCategoryTotals = true

	resp, err := req.Do()
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if !resp.Success {
		t.Error("Success = false, want true")
	}
	if len(resp.Data) != 1 {
		t.Fatalf("len(Data) = %d, want 1", len(resp.Data))
	}
	if resp.Data[0].Rowid != "pmt-1" {
		t.Errorf("Data[0].Rowid = %q, want %q", resp.Data[0].Rowid, "pmt-1")
	}
	if want := "/v1/payments/transaction_report.json"; gotPath != want {
		t.Errorf("request path = %q, want %q", gotPath, want)
	}
}

func TestGetAllRefundsRequestOffline(t *testing.T) {
	var gotPath, gotQuery string
	srv := httptest.NewServer(jsonHandler(&gotPath, &gotQuery,
		`{"success":true,"data":[{}]}`))
	defer srv.Close()

	c := newTestClient(t, srv)
	req := c.NewGetAllRefundsRequest()
	req.QueryParams().StartDate = tripleseat.Date{time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().EndDate = tripleseat.Date{time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().IncludeCategoryTotals = true

	resp, err := req.Do()
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if !resp.Success {
		t.Error("Success = false, want true")
	}
	if len(resp.Data) != 1 {
		t.Fatalf("len(Data) = %d, want 1", len(resp.Data))
	}
	if want := "/v1/payments/refund_report.json"; gotPath != want {
		t.Errorf("request path = %q, want %q", gotPath, want)
	}
}
