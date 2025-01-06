package tripleseat_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetBookingByIdRequest(t *testing.T) {
	req := client.NewGetBookingByIdRequest()
	req.PathParams().BookingID = 42946736
	req.QueryParams().ShowFinancials = true
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
