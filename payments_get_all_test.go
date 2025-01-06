package tripleseat_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	tripleseat "github.com/omniboost/go-tripleseat"
)

func TestGetAllPaymentsRequest(t *testing.T) {
	req := client.NewGetAllPaymentsRequest()
	req.QueryParams().StartDate = tripleseat.Date{time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().EndDate = tripleseat.Date{time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().IncludeCategoryTotals = true
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
