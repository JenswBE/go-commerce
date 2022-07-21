//go:build e2e
// +build e2e

package e2e

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/config"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const BaseURL = "http://127.0.0.1:8090/api/"

type E2ETestSuite struct {
	suite.Suite
	publicClient *openapi.APIClient // Unauthenticated client
	authClient   *openapi.APIClient // Authenticated client
}

func Test_E2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}

func (s *E2ETestSuite) SetupSuite() {
	// Parse config
	apiConfig, err := config.ParseConfig()
	require.NoError(s.T(), err)

	// Create public client
	s.publicClient = openapi.NewAPIClient(newE2EConfig(*apiConfig))

	// Setup authenticated client
	authConfig := newE2EConfig(*apiConfig)
	switch apiConfig.Authentication.Type {
	case config.AuthTypeBasicAuth:
		creds := fmt.Sprintf("%s:%s", apiConfig.Authentication.BasicAuth.Username, apiConfig.Authentication.BasicAuth.Password)
		fmt.Println(creds)
		authConfig.AddDefaultHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(creds)))
	case config.AuthTypeOIDC:
		// Get token
		formData := url.Values{}
		formData.Add("grant_type", "password")
		formData.Add("client_id", "go-commerce-admin")
		formData.Add("username", "e2e")
		formData.Add("password", "e2e")
		res, err := http.PostForm("http://127.0.0.1:9001/realms/go-commerce/protocol/openid-connect/token", formData)
		require.NoError(s.T(), err)

		// Parse token
		type tokenResponse struct {
			AccessToken string `json:"access_token"`
		}
		var tokenData tokenResponse
		err = json.NewDecoder(res.Body).Decode(&tokenData)
		require.NoError(s.T(), err)
		require.NotEmpty(s.T(), tokenData.AccessToken, "A valid access token should have been returned")
		authConfig.AddDefaultHeader("Authorization", "Bearer "+tokenData.AccessToken)
	}
	s.authClient = openapi.NewAPIClient(authConfig)
}

func newE2EConfig(apiConfig config.Config) *openapi.Configuration {
	config := openapi.NewConfiguration()
	config.Scheme = "http"
	config.Host = fmt.Sprintf("127.0.0.1:%d", apiConfig.Server.Port)
	config.Servers[0].URL = ""
	return config
}

func extractHTTPBody(t *testing.T, r *http.Response) string {
	if r == nil {
		return "response is nil"
	}

	body, err := io.ReadAll(r.Body)
	require.NoError(t, err)
	return string(body)
}
