package config

import (
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/stretchr/testify/require"
)

func Test_listContent_Success(t *testing.T) {
	// Setup test
	c, r := handler.SetupGinTest(t, "GET", "", nil, nil)
	configHandler := setupHandlerTest()

	// Call handler
	configHandler.getConfig(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ConfigOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)
}
