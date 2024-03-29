package product

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/JenswBE/go-commerce/utils/generics"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
)

func Test_listCategories_Success(t *testing.T) {
	// Setup test
	c, r := handler.SetupGinTest(t, "GET", "", nil, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("ListCategories", mock.Anything).Return(fixtures.CategorySlice(), nil)

	// Call handler
	productHandler.listCategories(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.CategoryListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "ListCategories", map[string]imageproxy.ImageConfig(nil))
}

func Test_getCategory_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.CategoryID}}
	c, r := handler.SetupGinTest(t, "GET", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetCategory", mock.Anything, mock.Anything).Return(fixtures.Category(), nil)

	// Call handler
	productHandler.getCategory(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.CategoryOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetCategory", generics.Must(entities.NewIDFromString(fixtures.CategoryID)), map[string]imageproxy.ImageConfig(nil))
}
