//go:build e2e

package e2e

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tebeka/selenium"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/JenswBE/go-commerce/admin"
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/config"
	"github.com/JenswBE/go-commerce/usecases"
)

type E2ETestSuite struct {
	suite.Suite
	config    *config.Config
	db        *gorm.DB
	apiClient *openapi.APIClient
	swd       selenium.WebDriver
}

func Test_E2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}

func (s *E2ETestSuite) SetupSuite() {
	// Parse config
	svcConfig, err := config.ParseConfig()
	require.NoError(s.T(), err)
	s.config = svcConfig

	// Update config
	svcConfig.Authentication.Type = config.AuthTypeNone
	svcConfig.Server.Port = 9999
	svcConfig.Database.Default.Port = 5433
	svcConfig.Database.Default.Database = "e2e"
	svcConfig.Database.Default.User = "e2e"
	svcConfig.Database.Default.Password = "e2e"

	// Connect to DB
	e2eDSN := config.BuildDSN(svcConfig.Database.Default, svcConfig.Database.Default)
	s.db, err = gorm.Open(postgres.Open(e2eDSN), &gorm.Config{})
	if err != nil {
		require.NoError(s.T(), err)
	}

	// Start GoCommerce service
	go usecases.StartService(svcConfig)

	// Create API client
	s.apiClient = openapi.NewAPIClient(newAPIConfig(*svcConfig))

	// Connect to Selenium
	caps := selenium.Capabilities{"browserName": "firefox"}
	s.swd, err = selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 4444))
	if err != nil {
		require.NoError(s.T(), err)
	}
}

func (s *E2ETestSuite) SetupTest() {
	for _, table := range []string{
		"categories",
		"contents",
		"events",
		"manufacturers",
		"products",
	} {
		require.NoError(s.T(), s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", table)).Error)
	}
}

func (s *E2ETestSuite) TearDownSuite() {
	lo.Must0(s.swd.Quit())
}

func newAPIConfig(apiConfig config.Config) *openapi.Configuration {
	config := openapi.NewConfiguration()
	config.Scheme = "http"
	config.Host = fmt.Sprintf("127.0.0.1:%d", apiConfig.Server.Port)
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

func (s *E2ETestSuite) rootURL(pathParts ...string) string {
	host := fmt.Sprintf("http://host.containers.internal:%d/", s.config.Server.Port)
	return host + path.Join(pathParts...) + "/"
}

func (s *E2ETestSuite) adminURL(pathParts ...string) string {
	pathParts = append([]string{strings.TrimPrefix(admin.PrefixAdmin, "/")}, pathParts...)
	return s.rootURL(pathParts...)
}

func must(s *E2ETestSuite, err error) {
	require.NoError(s.T(), err)
}

func must1[T any](s *E2ETestSuite, result T, err error) T {
	require.NoError(s.T(), err)
	return result
}

func (s *E2ETestSuite) swdMustGetAdmin(adminURL string) {
	lo.Must0(s.swd.Get(s.adminURL(adminURL)))
}

func (s *E2ETestSuite) swdMustFindElement(by, value string) selenium.WebElement {
	elem, err := s.swd.FindElement(by, value)
	require.NoError(s.T(), err)
	return elem
}

func (s *E2ETestSuite) swdMustFindElements(by, value string) []selenium.WebElement {
	elems, err := s.swd.FindElements(by, value)
	require.NoError(s.T(), err)
	return elems
}
