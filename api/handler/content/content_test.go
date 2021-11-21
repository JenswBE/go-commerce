package content

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_listContent_Success(t *testing.T) {
	// Setup test
	c, r := handler.SetupGinTest(t, "GET", "", nil, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("ListContent", mock.Anything).Return(fixtures.ContentSlice(), nil)

	// Call handler
	productHandler.listContent(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ContentListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "ListContent")
}

func Test_getContent_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "content_name", Value: fixtures.Content().Name}}
	c, r := handler.SetupGinTest(t, "GET", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetContent", mock.Anything, mock.Anything).Return(fixtures.Content(), nil)

	// Call handler
	productHandler.getContent(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ContentOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetContent", fixtures.Content().Name)
}

func Test_updateContent_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "content_name", Value: fixtures.Content().Name}}
	body, err := json.Marshal(fixtures.ContentOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "PUT", "", params, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateContent", mock.Anything).Return(fixtures.Content(), nil)

	// Call handler
	productHandler.updateContent(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ContentOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpdateContent", mock.MatchedBy(func(actual *entities.Content) bool {
		require.Equal(t, fixtures.Content(), actual)
		return true
	}))
}
