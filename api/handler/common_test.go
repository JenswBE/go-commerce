package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/utils/shortid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// ###########################
// #          TESTS          #
// ###########################

func Test_parseIDParam_Success(t *testing.T) {
	// Setup test
	presenter := presenter.New(shortid.NewFakeService())
	value := uuid.New()
	c, _ := setupGinTest("", "", gin.Params{{Key: "test_id", Value: value.String()}}, "")

	// Call helper
	result, success := parseIDParam(c, "test_id", presenter)

	// Assert results
	require.Equal(t, value, result)
	require.True(t, success)
}

func Test_parseIDParam_ParamNotProvided_Failure(t *testing.T) {
	// Setup test
	presenter := presenter.New(shortid.NewFakeService())
	c, w := setupGinTest("", "", nil, "")

	// Call helper
	result, success := parseIDParam(c, "test_id", presenter)

	// Assert results
	require.Contains(t, w.Body.String(), "mandatory")
	require.Equal(t, 400, w.Code)
	require.Equal(t, uuid.Nil, result)
	require.False(t, success)
}

// ###########################
// #         HELPERS         #
// ###########################

func setupGinTest(method, path string, params gin.Params, body string) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	bodyReader := strings.NewReader(body)
	c.Request, _ = http.NewRequest(method, path, bodyReader)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
	return c, w
}
