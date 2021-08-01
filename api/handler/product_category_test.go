package handler

import (
	"encoding/json"
	"net/http"
	"testing"

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
	c, r := setupGinTest("", "", nil, nil)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("ListCategories", mock.Anything).Return(fixtures.CategorySlice(), nil)

	// Call handler
	handler.listCategories(c)

	// Assert result
	requireEqualJSON(t, fixtures.CategoryOpenAPISlice(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "ListCategories", *new(*imageproxy.ImageConfig))
}

func Test_getCategory_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.CategoryID}}
	c, r := setupGinTest("", "", params, nil)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetCategory", mock.Anything, mock.Anything).Return(fixtures.Category(), nil)

	// Call handler
	handler.getCategory(c)

	// Assert result
	requireEqualJSON(t, fixtures.CategoryOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetCategory", uuid.MustParse(fixtures.CategoryID), *new(*imageproxy.ImageConfig))
}

func Test_createCategory_Success(t *testing.T) {
	// Setup test
	body, err := json.Marshal(fixtures.CategoryOpenAPI())
	require.NoError(t, err)
	c, r := setupGinTest("", "", nil, body)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("CreateCategory", mock.Anything).Return(fixtures.Category(), nil)

	// Call handler
	handler.createCategory(c)

	// Assert result
	requireEqualJSON(t, fixtures.CategoryOpenAPI(), r)
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
	c, r := setupGinTest("", "", params, body)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateCategory", mock.Anything).Return(fixtures.Category(), nil)

	// Call handler
	handler.updateCategory(c)

	// Assert result
	requireEqualJSON(t, fixtures.CategoryOpenAPI(), r)
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
	c, r := setupGinTest("", "", params, nil)
	handler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteCategory", mock.Anything).Return(nil)

	// Call handler
	handler.deleteCategory(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteCategory", uuid.MustParse(fixtures.CategoryID))
}
