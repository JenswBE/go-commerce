package product

import (
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_listProducts_Success(t *testing.T) {
	// Setup test
	c, r := handler.SetupGinTest(t, "GET", "", nil, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("ListProducts", mock.Anything).Return(fixtures.ProductSlice(), nil)

	// Call handler
	productHandler.listProducts(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ProductListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "ListProducts", map[string]imageproxy.ImageConfig(nil))
}

func Test_getProduct_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	c, r := handler.SetupGinTest(t, "GET", "/?resolve=true", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetProduct", mock.Anything, mock.Anything, mock.Anything).Return(fixtures.ResolvedProduct(), nil)

	// Call handler
	productHandler.getProduct(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ResolvedProductOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetProduct", uuid.MustParse(fixtures.ProductID), true, map[string]imageproxy.ImageConfig(nil))
}
