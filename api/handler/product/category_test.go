package product

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
	usecaseMock.AssertCalled(t, "GetCategory", uuid.MustParse(fixtures.CategoryID), map[string]imageproxy.ImageConfig(nil))
}

func Test_createCategory_Success(t *testing.T) {
	// Setup test
	body, err := json.Marshal(fixtures.CategoryOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "POST", "", nil, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("CreateCategory", mock.Anything).Return(fixtures.Category(), nil)

	// Call handler
	productHandler.createCategory(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.CategoryOpenAPI(), r)
	require.Equal(t, http.StatusCreated, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "CreateCategory", mock.MatchedBy(func(cat *entities.Category) bool {
		expected := fixtures.Category()
		expected.ID = uuid.Nil
		expected.Image = nil
		require.Equal(t, expected, cat)
		return true
	}))
}

func Test_updateCategory_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.CategoryID}}
	body, err := json.Marshal(fixtures.CategoryOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "PUT", "", params, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateCategory", mock.Anything).Return(fixtures.Category(), nil)

	// Call handler
	productHandler.updateCategory(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.CategoryOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpdateCategory", mock.MatchedBy(func(cat *entities.Category) bool {
		expected := fixtures.Category()
		expected.Image = nil
		require.Equal(t, expected, cat)
		return true
	}))
}

func Test_deleteCategory_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.CategoryID}}
	c, r := handler.SetupGinTest(t, "DELETE", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteCategory", mock.Anything).Return(nil)

	// Call handler
	productHandler.deleteCategory(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteCategory", uuid.MustParse(fixtures.CategoryID))
}

func Test_upsertCategoryImage_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.CategoryID}}
	body, writer := fixtures.MultipartSingleFile()
	c, r := handler.SetupGinTest(t, "PUT", "?"+fixtures.ImageConfigQuery, params, body.Bytes())
	c.Request.Header.Set("Content-Type", writer.FormDataContentType())
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpsertCategoryImage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(fixtures.Category(), nil)

	// Call handler
	productHandler.upsertCategoryImage(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ImageOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpsertCategoryImage", uuid.MustParse(fixtures.CategoryID), "test.jpg", []byte("test-jpg"), fixtures.ImageConfigMap())
}

func Test_deleteCategoryImage_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.CategoryID}}
	c, r := handler.SetupGinTest(t, "DELETE", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteCategoryImage", mock.Anything).Return(nil)

	// Call handler
	productHandler.deleteCategoryImage(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteCategoryImage", uuid.MustParse(fixtures.CategoryID))
}
