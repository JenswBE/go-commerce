package content

import (
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_getContent_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "content_name", Value: fixtures.Content().Name}}
	c, r := handler.SetupGinTest(t, "GET", "", params, nil)
	contentHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetContent", mock.Anything, mock.Anything).Return(fixtures.Content(), nil)

	// Call handler
	contentHandler.getContent(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ContentOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetContent", fixtures.Content().Name)
}
