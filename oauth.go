package tripleseat

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Oauth2Config handles the OAuth 2.0 authorization code flow with refresh tokens.
type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	return &Oauth2Config{
		Config: oauth2.Config{
			Scopes: []string{},
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://login.tripleseat.com/oauth2/authorize",
				TokenURL:  "https://api.tripleseat.com/oauth2/token",
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
	}
}

func (c *Oauth2Config) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return c.Config.Client(ctx, token)
}

// Oauth2ClientCredentialsConfig handles the client credentials flow.
type Oauth2ClientCredentialsConfig struct {
	clientcredentials.Config
}

func NewOauth2ClientCredentialsConfig(clientID, clientSecret string) *Oauth2ClientCredentialsConfig {
	return &Oauth2ClientCredentialsConfig{
		Config: clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     "https://api.tripleseat.com/oauth/token",
			AuthStyle:    oauth2.AuthStyleInParams,
		},
	}
}

func (c *Oauth2ClientCredentialsConfig) Client(ctx context.Context) *http.Client {
	return c.Config.Client(ctx)
}
