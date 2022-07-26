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
