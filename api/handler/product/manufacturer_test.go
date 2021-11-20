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

func Test_listManufacturers_Success(t *testing.T) {
	// Setup test
	c, r := handler.SetupGinTest(t, "GET", "", nil, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("ListManufacturers", mock.Anything).Return(fixtures.ManufacturerSlice(), nil)

	// Call handler
	productHandler.listManufacturers(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ManufacturerListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "ListManufacturers", map[string]imageproxy.ImageConfig(nil))
}

func Test_getManufacturer_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ManufacturerID}}
	c, r := handler.SetupGinTest(t, "GET", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetManufacturer", mock.Anything, mock.Anything).Return(fixtures.Manufacturer(), nil)

	// Call handler
	productHandler.getManufacturer(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ManufacturerOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetManufacturer", uuid.MustParse(fixtures.ManufacturerID), map[string]imageproxy.ImageConfig(nil))
}

func Test_createManufacturer_Success(t *testing.T) {
	// Setup test
	body, err := json.Marshal(fixtures.ManufacturerOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "POST", "", nil, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("CreateManufacturer", mock.Anything).Return(fixtures.Manufacturer(), nil)

	// Call handler
	productHandler.createManufacturer(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ManufacturerOpenAPI(), r)
	require.Equal(t, http.StatusCreated, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "CreateManufacturer", mock.MatchedBy(func(cat *entities.Manufacturer) bool {
		expected := fixtures.Manufacturer()
		expected.ID = uuid.Nil
		expected.Image = nil
		require.Equal(t, expected, cat)
		return true
	}))
}

func Test_updateManufacturer_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ManufacturerID}}
	body, err := json.Marshal(fixtures.ManufacturerOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "PUT", "", params, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateManufacturer", mock.Anything).Return(fixtures.Manufacturer(), nil)

	// Call handler
	productHandler.updateManufacturer(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ManufacturerOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpdateManufacturer", mock.MatchedBy(func(cat *entities.Manufacturer) bool {
		expected := fixtures.Manufacturer()
		expected.Image = nil
		require.Equal(t, expected, cat)
		return true
	}))
}

func Test_deleteManufacturer_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ManufacturerID}}
	c, r := handler.SetupGinTest(t, "DELETE", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteManufacturer", mock.Anything).Return(nil)

	// Call handler
	productHandler.deleteManufacturer(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteManufacturer", uuid.MustParse(fixtures.ManufacturerID))
}

func Test_upsertManufacturerImage_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ManufacturerID}}
	body, writer := fixtures.MultipartSingleFile()
	c, r := handler.SetupGinTest(t, "PUT", "?"+fixtures.ImageConfigQuery, params, body.Bytes())
	c.Request.Header.Set("Content-Type", writer.FormDataContentType())
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpsertManufacturerImage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(fixtures.Manufacturer(), nil)

	// Call handler
	productHandler.upsertManufacturerImage(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.ImageOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpsertManufacturerImage", uuid.MustParse(fixtures.ManufacturerID), "test.jpg", []byte("test-jpg"), fixtures.ImageConfigMap())
}

func Test_deleteManufacturerImage_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.ManufacturerID}}
	c, r := handler.SetupGinTest(t, "DELETE", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteManufacturerImage", mock.Anything).Return(nil)

	// Call handler
	productHandler.deleteManufacturerImage(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteManufacturerImage", uuid.MustParse(fixtures.ManufacturerID))
}
