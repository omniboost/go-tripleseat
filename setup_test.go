package tripleseat_test

import (
	"context"
	"log"
	"os"
	"testing"

	tripleseat "github.com/omniboost/go-tripleseat"
)

var (
	client *tripleseat.Client
)

func TestMain(m *testing.M) {
	clientID := os.Getenv("TRIPLESEAT_CLIENT_ID")
	clientSecret := os.Getenv("TRIPLESEAT_CLIENT_SECRET")
	tokenURL := os.Getenv("TRIPLESEAT_TOKEN_URL")

	oauthConfig := tripleseat.NewOauth2Config(clientID, clientSecret)

	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	httpClient, err := oauthConfig.Client(context.Background())
	if err != nil {
		log.Fatalf("failed to create http client: %v", err)
	}

	client = tripleseat.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	m.Run()
}
