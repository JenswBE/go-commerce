package product

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/entities"
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

func Test_createProduct_Success(t *testing.T) {
	// Setup test
	body, err := json.Marshal(fixtures.ProductOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "POST", "", nil, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("CreateProduct", mock.Anything).Return(fixtures.Product(), nil)

	// Call handler
	productHandler.createProduct(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ProductOpenAPI(), r)
	require.Equal(t, http.StatusCreated, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "CreateProduct", mock.MatchedBy(func(actual *entities.Product) bool {
		expected := fixtures.Product()
		expected.ID = uuid.Nil
		expected.CreatedAt = time.Time{}
		expected.UpdatedAt = time.Time{}
		expected.Images = nil
		require.Equal(t, expected, actual)
		return true
	}))
}

func Test_updateProduct_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	body, err := json.Marshal(fixtures.ProductOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "PUT", "", params, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateProduct", mock.Anything).Return(fixtures.Product(), nil)

	// Call handler
	productHandler.updateProduct(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ProductOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpdateProduct", mock.MatchedBy(func(actual *entities.Product) bool {
		expected := fixtures.Product()
		expected.CreatedAt = time.Time{}
		expected.UpdatedAt = time.Time{}
		expected.Images = nil
		require.Equal(t, expected, actual)
		return true
	}))
}

func Test_deleteProduct_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	c, r := handler.SetupGinTest(t, "DELETE", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteProduct", mock.Anything).Return(nil)

	// Call handler
	productHandler.deleteProduct(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteProduct", uuid.MustParse(fixtures.ProductID))
}

func Test_listProductImages_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	c, r := handler.SetupGinTest(t, "GET", "?"+fixtures.ImageConfigQuery, params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetProduct", mock.Anything, mock.Anything, mock.Anything).Return(fixtures.ResolvedProduct(), nil)

	// Call handler
	productHandler.listProductImages(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ImageListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetProduct", uuid.MustParse(fixtures.ProductID), false, fixtures.ImageConfigMap())
}

func Test_addProductImages_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	body, writer := fixtures.MultipartMultipleFiles()
	c, r := handler.SetupGinTest(t, "POST", "?"+fixtures.ImageConfigQuery, params, body.Bytes())
	c.Request.Header.Set("Content-Type", writer.FormDataContentType())
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("AddProductImages", mock.Anything, mock.Anything, mock.Anything).Return(fixtures.Product(), nil)

	// Call handler
	productHandler.addProductImages(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ImageListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "AddProductImages", uuid.MustParse(fixtures.ProductID), fixtures.MultipartMultipleFilesMap(), fixtures.ImageConfigMap())
}

func Test_updateProductImage_Success(t *testing.T) {
	// Setup test
	params := gin.Params{
		{Key: "id", Value: fixtures.ProductID},
		{Key: "image_id", Value: fixtures.ImageID},
	}
	body, err := json.Marshal(openapi.Image{Order: 5})
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "PUT", "", params, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateProductImage", mock.Anything, mock.Anything, mock.Anything).Return(fixtures.ImageSlice(), nil)

	// Call handler
	productHandler.updateProductImage(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ImageOpenAPISlice(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpdateProductImage", uuid.MustParse(fixtures.ProductID), uuid.MustParse(fixtures.ImageID), 5)
}

func Test_deleteProductImage_Success(t *testing.T) {
	// Setup test
	params := gin.Params{
		{Key: "id", Value: fixtures.ProductID},
		{Key: "image_id", Value: fixtures.ImageID},
	}
	c, r := handler.SetupGinTest(t, "DELETE", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteProductImage", mock.Anything, mock.Anything).Return(nil)

	// Call handler
	productHandler.deleteProductImage(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteProductImage", uuid.MustParse(fixtures.ProductID), uuid.MustParse(fixtures.ImageID))
}
