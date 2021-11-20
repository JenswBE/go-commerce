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
	"github.com/JenswBE/go-commerce/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

// ParseIDParam tries to parse parameter with the given name as an UUID or short ID.
// On failure, an error is set on the Gin context.
//
//   p := presenter.New()
//   id, ok := ParseIDParam(c, "id", p)
//   if !ok {
// 	   return // Response already set on Gin context
//   }
func ParseIDParam(c *gin.Context, name string, p *presenter.Presenter) (uuid.UUID, bool) {
	// Parse param
	pID, ok := c.Params.Get(name)
	if !ok {
		err := entities.NewError(400, openapi.GOCOMERRORCODE_PARAMETER_MISSING, name, nil)
		c.JSON(ErrToResponse(err))
		return uuid.Nil, false
	}

	// Parse ID
	id, err := p.ParseID(pID)
	if err != nil {
		c.JSON(ErrToResponse(err))
		return uuid.Nil, false
	}

	// Parse successful
	return id, true
}

// ErrToResponse checks if the provided error is a GoComError.
// If yes, status and embedded error message are returned.
// If no, status is 500 and provided error message are returned.
func ErrToResponse(e error) (int, *entities.GoComError) {
	if err, ok := e.(*entities.GoComError); ok {
		return err.Status, err
	}
	log.Warn().Err(e).Stringer("error_type", reflect.TypeOf(e)).Msg("API received an non-GoComError error")
	return 500, entities.NewError(500, openapi.GOCOMERRORCODE_UNKNOWN_ERROR, "", e).(*entities.GoComError)
}

// ###########################
// #       TEST HELPERS      #
// ###########################

func SetupGinTest(t *testing.T, method, path string, params gin.Params, body []byte) (*gin.Context, *httptest.ResponseRecorder) {
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

// RequireEqualJSON unmarshals the body from the provided recorder to the same type as expected.
// Next, it asserts this result against the expected value.
func RequireEqualJSON(t *testing.T, expected interface{}, recorder *httptest.ResponseRecorder) {
	actual := reflect.New(reflect.TypeOf(expected))
	body := recorder.Body.Bytes()
	err := json.Unmarshal(body, actual.Interface())
	require.NoErrorf(t, err, `Response body: %s`, string(body))
	// Has better support for time.Time than require.Equal
	// See https://github.com/stretchr/testify/issues/1078 for more info
	require.Empty(t, cmp.Diff(expected, reflect.Indirect(actual).Interface()))
}
