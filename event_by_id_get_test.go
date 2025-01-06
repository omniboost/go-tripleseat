package tripleseat_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetEventByIdRequest(t *testing.T) {
	req := client.NewGetEventByIdRequest()
	req.PathParams().EventID = 45103683
	req.QueryParams().ShowFinancials = true
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
