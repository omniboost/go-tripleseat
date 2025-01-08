package tripleseat

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scope                = ""
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	oauth2.Config
	ClientCredentialsConfig clientcredentials.Config
}

func NewOauth2Config(clientID, clientSecret string) *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://api.tripleseat.com/oauth/token",
				TokenURL: "https://api.tripleseat.com/oauth/token",
			},
		},
		ClientCredentialsConfig: clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     "https://api.tripleseat.com/oauth/token",
		},
	}

	config.SetBaseURL(&BaseURL)
	return config
}

func (c *Oauth2Config) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	// These are not registered in the oauth library by default
	oauth2.RegisterBrokenAuthHeaderProvider(baseURL.String())

	c.Config.Endpoint = oauth2.Endpoint{
		AuthURL:  baseURL.String() + "/oauth/token",
		TokenURL: baseURL.String() + "/oauth/token",
	}
	c.ClientCredentialsConfig.TokenURL = baseURL.String() + "/oauth/token"
}

func (c *Oauth2Config) GetToken(ctx context.Context) (*oauth2.Token, error) {
	token, err := c.ClientCredentialsConfig.Token(ctx)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (c *Oauth2Config) Client(ctx context.Context) (*http.Client, error) {
	token, err := c.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	return c.Config.Client(ctx, token), nil
}
