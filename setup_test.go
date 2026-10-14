package tripleseat_test

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	tripleseat "github.com/omniboost/go-tripleseat"
	"golang.org/x/oauth2"
)

var (
	client *tripleseat.Client
)

func TestMain(m *testing.M) {
	authType := os.Getenv("TRIPLESEAT_AUTH_TYPE")
	clientID := os.Getenv("TRIPLESEAT_CLIENT_ID")
	clientSecret := os.Getenv("TRIPLESEAT_CLIENT_SECRET")
	tokenURL := os.Getenv("TRIPLESEAT_TOKEN_URL")
	refreshToken := os.Getenv("TRIPLESEAT_REFRESH_TOKEN")

	// Without credentials, leave the shared client nil so the live API tests
	// skip (see requireLiveClient) instead of making real network calls. The
	// offline httptest-based tests build their own clients and do not use this.
	if clientID == "" {
		os.Exit(m.Run())
	}

	var httpClient *http.Client

	switch authType {
	case "", "oauth2":
		oauthConfig := tripleseat.NewOauth2Config()
		oauthConfig.ClientID = clientID
		oauthConfig.ClientSecret = clientSecret

		if tokenURL != "" {
			oauthConfig.Endpoint.TokenURL = tokenURL
		}

		httpClient = oauthConfig.Client(context.Background(), &oauth2.Token{
			RefreshToken: refreshToken,
		})

	case "client_credentials":
		oauthConfig := tripleseat.NewOauth2ClientCredentialsConfig(clientID, clientSecret)
		if tokenURL != "" {
			oauthConfig.TokenURL = tokenURL
		}
		httpClient = oauthConfig.Client(context.Background())

	default:
		log.Fatalf("unknown TRIPLESEAT_AUTH_TYPE: %s", authType)
	}

	client = tripleseat.NewClient(httpClient)
	client.SetDisallowUnknownFields(true)

	os.Exit(m.Run())
}

// requireLiveClient skips a test unless live Tripleseat credentials were
// provided via the TRIPLESEAT_* environment variables. This keeps `go test`
// offline and green by default; the live tests print real API responses and
// should only run against an authorized account.
func requireLiveClient(t *testing.T) {
	t.Helper()
	if client == nil {
		t.Skip("set TRIPLESEAT_CLIENT_ID (and related TRIPLESEAT_* vars) to run live API tests")
	}
}
