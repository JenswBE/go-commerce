//go:build e2e

package e2e

import (
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func (s *E2ETestSuite) TestBasicValidations() {
	// 1. Validate "/api" redirects to "/api/docs/"
	// 2. Validate the API docs are shown
	lo.Must0(s.swd.Get(s.rootURL("api")))
	require.Equal(s.T(), s.rootURL("api/docs"), lo.Must(s.swd.CurrentURL()))
	require.Contains(s.T(), lo.Must(s.swd.Title()), "Swagger UI")
}
