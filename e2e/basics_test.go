//go:build e2e

package e2e

import (
	"github.com/stretchr/testify/require"
)

func (s *E2ETestSuite) TestBasicValidations() {
	// 1. Validate "/api" redirects to "/api/docs/"
	// 2. Validate the API docs are shown
	s.must(s.swd.Get(s.rootURL("api")))
	currentURL, err := s.swd.CurrentURL()
	s.must(err)
	require.Equal(s.T(), s.rootURL("api/docs"), currentURL)
	currentTitle, err := s.swd.Title()
	s.must(err)
	require.Contains(s.T(), currentTitle, "Swagger UI")
}
