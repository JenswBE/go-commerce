package product

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/fixtures"
	mocks "github.com/JenswBE/go-commerce/mocks/usecases/product"
	"github.com/JenswBE/go-commerce/utils/sanitizer"
	"github.com/JenswBE/go-commerce/utils/shortid"
)

func Test_parseImageConfigsParam_AllParamsSet_Success(t *testing.T) {
	// Setup test
	c, _ := handler.SetupGinTest(t, "", "?"+fixtures.ImageConfigQuery, nil, nil)

	// Call function
	result, err := parseImageConfigsParam(c)

	// Assert results
	require.NoError(t, err)
	require.Equal(t, fixtures.ImageConfigMap(), result)
}

func Test_parseImageConfigsParam_NoParamsSet_Success(t *testing.T) {
	// Setup test
	c, _ := handler.SetupGinTest(t, "", "", nil, nil)

	// Call function
	result, err := parseImageConfigsParam(c)

	// Assert results
	require.NoError(t, err)
	require.Nil(t, result)
}

func setupHandlerTest() (*ProductHandler, *mocks.Usecase) {
	presenter := presenter.New(shortid.NewFakeService(), sanitizer.NewFakeService())
	usecase := &mocks.Usecase{}
	handler := NewProductHandler(presenter, usecase)
	return handler, usecase
}
