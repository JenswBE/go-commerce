package handler

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

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
	c, r := setupGinTest(t, "GET", "", nil, nil)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("ListProducts", mock.Anything).Return(fixtures.ProductSlice(), nil)

	// Call handler
	handler.listProducts(c)

	// Assert result
	requireEqualJSON(t, fixtures.ProductListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "ListProducts", *new(*imageproxy.ImageConfig))
}

func Test_getProduct_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	c, r := setupGinTest(t, "GET", "", params, nil)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetProduct", mock.Anything, mock.Anything).Return(fixtures.Product(), nil)

	// Call handler
	handler.getProduct(c)

	// Assert result
	requireEqualJSON(t, fixtures.ProductOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetProduct", uuid.MustParse(fixtures.ProductID), *new(*imageproxy.ImageConfig))
}

func Test_createProduct_Success(t *testing.T) {
	// Setup test
	body, err := json.Marshal(fixtures.ProductOpenAPI())
	require.NoError(t, err)
	c, r := setupGinTest(t, "POST", "", nil, body)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("CreateProduct", mock.Anything).Return(fixtures.Product(), nil)

	// Call handler
	handler.createProduct(c)

	// Assert result
	requireEqualJSON(t, fixtures.ProductOpenAPI(), r)
	require.Equal(t, http.StatusCreated, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "CreateProduct", mock.MatchedBy(func(product *entities.Product) bool {
		expected := fixtures.Product()
		expected.ID = uuid.Nil
		expected.CreatedAt = time.Time{}
		expected.UpdatedAt = time.Time{}
		expected.Images = nil
		require.Equal(t, expected, product)
		return true
	}))
}

func Test_updateProduct_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	body, err := json.Marshal(fixtures.ProductOpenAPI())
	require.NoError(t, err)
	c, r := setupGinTest(t, "PUT", "", params, body)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateProduct", mock.Anything).Return(fixtures.Product(), nil)

	// Call handler
	handler.updateProduct(c)

	// Assert result
	requireEqualJSON(t, fixtures.ProductOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpdateProduct", mock.MatchedBy(func(product *entities.Product) bool {
		expected := fixtures.Product()
		expected.CreatedAt = time.Time{}
		expected.UpdatedAt = time.Time{}
		expected.Images = nil
		require.Equal(t, expected, product)
		return true
	}))
}

func Test_deleteProduct_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	c, r := setupGinTest(t, "DELETE", "", params, nil)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteProduct", mock.Anything).Return(nil)

	// Call handler
	handler.deleteProduct(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteProduct", uuid.MustParse(fixtures.ProductID))
}

func Test_listProductImages_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	c, r := setupGinTest(t, "GET", "?"+fixtures.ImageConfigQuery, params, nil)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetProduct", mock.Anything, mock.Anything).Return(fixtures.Product(), nil)

	// Call handler
	handler.listProductImages(c)

	// Assert result
	requireEqualJSON(t, fixtures.ImageListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetProduct", uuid.MustParse(fixtures.ProductID), fixtures.ImageConfig())
}

func Test_addProductImages_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ProductID}}
	body, writer := fixtures.MultipartMultipleFiles()
	c, r := setupGinTest(t, "POST", "?"+fixtures.ImageConfigQuery, params, body.Bytes())
	c.Request.Header.Set("Content-Type", writer.FormDataContentType())
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("AddProductImages", mock.Anything, mock.Anything, mock.Anything).Return(fixtures.Product(), nil)

	// Call handler
	handler.addProductImages(c)

	// Assert result
	requireEqualJSON(t, fixtures.ImageListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "AddProductImages", uuid.MustParse(fixtures.ProductID), fixtures.MultipartMultipleFilesMap(), fixtures.ImageConfig())
}

func Test_updateProductImage_Success(t *testing.T) {
	// Setup test
	params := gin.Params{
		{Key: "id", Value: fixtures.ProductID},
		{Key: "image_id", Value: fixtures.ImageID},
	}
	body, err := json.Marshal(openapi.Image{Order: 5})
	require.NoError(t, err)
	c, r := setupGinTest(t, "PUT", "", params, body)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateProductImage", mock.Anything, mock.Anything, mock.Anything).Return(fixtures.ImageSlice(), nil)

	// Call handler
	handler.updateProductImage(c)

	// Assert result
	requireEqualJSON(t, fixtures.ImageOpenAPISlice(), r)
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
	c, r := setupGinTest(t, "DELETE", "", params, nil)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteProductImage", mock.Anything, mock.Anything).Return(nil)

	// Call handler
	handler.deleteProductImage(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteProductImage", uuid.MustParse(fixtures.ProductID), uuid.MustParse(fixtures.ImageID))
}
