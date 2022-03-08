// +build e2e

package e2e

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const BaseURL = "http://localhost:8090/api/"

type E2ETestSuite struct {
	suite.Suite
	publicClient *openapi.APIClient // Unauthenticated client
	authClient   *openapi.APIClient // Authenticated client
}

func Test_E2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}

func (s *E2ETestSuite) SetupSuite() {
	// Get token
	formData := url.Values{}
	formData.Add("grant_type", "password")
	formData.Add("client_id", "go-commerce-admin")
	formData.Add("username", "e2e")
	formData.Add("password", "e2e")
	res, err := http.PostForm("http://127.0.0.1:9001/auth/realms/go-commerce/protocol/openid-connect/token", formData)
	require.NoError(s.T(), err)

	// Parse token
	type tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	var tokenData tokenResponse
	err = json.NewDecoder(res.Body).Decode(&tokenData)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), tokenData.AccessToken)

	// Create client
	s.publicClient = openapi.NewAPIClient(newE2EConfig())
	authConfig := newE2EConfig()
	log.Println(tokenData.AccessToken)
	authConfig.AddDefaultHeader("Authorization", "Bearer "+tokenData.AccessToken)
	s.authClient = openapi.NewAPIClient(authConfig)
}

func newE2EConfig() *openapi.Configuration {
	config := openapi.NewConfiguration()
	config.Scheme = "http"
	config.Host = "127.0.0.1:8090"
	return config
}

func extractHTTPBody(t *testing.T, r *http.Response) string {
	body, err := io.ReadAll(r.Body)
	require.NoError(t, err)
	return string(body)
}
