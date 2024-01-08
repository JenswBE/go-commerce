//go:build e2e

package e2e

import (
	"github.com/samber/lo"

	"github.com/JenswBE/go-commerce/utils/shortid"
)

func (s *E2ETestSuite) TestBasicValidations() {
	// 1. Validate "/api" redirects to "/api/docs/"
	// 2. Validate the API docs are shown
	lo.Must0(s.swd.Get(s.rootURL("api")))
	s.Require().Equal(s.rootURL("api/docs"), lo.Must(s.swd.CurrentURL()))
	s.Require().Contains(lo.Must(s.swd.Title()), "Swagger UI")
}

func decodeBase58UUID(input string) (string, error) {
	id, err := shortid.NewBase58Service().Decode(input)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
