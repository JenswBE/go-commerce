package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/JenswBE/go-commerce/api/openapi"
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/utils/sanitizer"
	"github.com/JenswBE/go-commerce/utils/shortid"
)

// ###########################
// #          TESTS          #
// ###########################

func Test_parseIDParam_Success(t *testing.T) {
	// Setup test
	presenter := presenter.New(shortid.NewFakeService(), sanitizer.NewFakeService())
	value := entities.NewID()
	c, _ := SetupGinTest(t, "", "", gin.Params{{Key: "test_id", Value: value.String()}}, nil)

	// Call function
	result, success := ParseIDParam(c, "test_id", presenter)

	// Assert results
	require.Equal(t, value, result)
	require.True(t, success)
}

func Test_parseIDParam_ParamNotProvided_Failure(t *testing.T) {
	// Setup test
	presenter := presenter.New(shortid.NewFakeService(), sanitizer.NewFakeService())
	c, w := SetupGinTest(t, "", "", nil, nil)

	// Call function
	result, success := ParseIDParam(c, "test_id", presenter)

	// Assert results
	require.Contains(t, w.Body.String(), openapi.GOCOMERRORCODE_PARAMETER_MISSING)
	require.Equal(t, 400, w.Code)
	require.True(t, result.IsNil(), "Result must be a nil UUID")
	require.False(t, success)
}
