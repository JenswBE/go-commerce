package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	mocks "github.com/JenswBE/go-commerce/mocks/usecases/product"
	"github.com/JenswBE/go-commerce/utils/shortid"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
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
	c, _ := setupGinTest(t, "", "", gin.Params{{Key: "test_id", Value: value.String()}}, nil)

	// Call function
	result, success := parseIDParam(c, "test_id", presenter)

	// Assert results
	require.Equal(t, value, result)
	require.True(t, success)
}

func Test_parseIDParam_ParamNotProvided_Failure(t *testing.T) {
	// Setup test
	presenter := presenter.New(shortid.NewFakeService())
	c, w := setupGinTest(t, "", "", nil, nil)

	// Call function
	result, success := parseIDParam(c, "test_id", presenter)

	// Assert results
	require.Contains(t, w.Body.String(), openapi.GOCOMERRORCODE_PARAMETER_MISSING)
	require.Equal(t, 400, w.Code)
	require.Equal(t, uuid.Nil, result)
	require.False(t, success)
}

// ###########################
// #         HELPERS         #
// ###########################

func setupGinTest(t *testing.T, method, path string, params gin.Params, body []byte) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	bodyReader := bytes.NewReader(body)
	var err error
	c.Request, err = http.NewRequest(method, path, bodyReader)
	require.NoError(t, err)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
	return c, w
}

func setupHandlerTest() (*ProductHandler, *mocks.Usecase) {
	presenter := presenter.New(shortid.NewFakeService())
	usecase := &mocks.Usecase{}
	handler := NewProductHandler(presenter, usecase)
	return handler, usecase
}

// requireEqualJSON unmarshals the body from the provided recorder to the same type as expected.
// Next, it asserts this result against the expected value.
func requireEqualJSON(t *testing.T, expected interface{}, recorder *httptest.ResponseRecorder) {
	actual := reflect.New(reflect.TypeOf(expected))
	body := recorder.Body.Bytes()
	err := json.Unmarshal(body, actual.Interface())
	require.NoErrorf(t, err, `Response body: %s`, string(body))
	// Has better support for time.Time than require.Equal
	// See https://github.com/stretchr/testify/issues/1078 for more info
	require.Empty(t, cmp.Diff(expected, reflect.Indirect(actual).Interface()))
}
