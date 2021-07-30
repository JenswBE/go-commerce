package handler

import (
	"testing"

	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/stretchr/testify/require"
)

func Test_parseImageConfigParams_AllParamsSet_Success(t *testing.T) {
	// Setup test
	c, _ := setupGinTest("", "/?img_w=300&img_h=200&img_r=fit", nil, "")

	// Call function
	result, err := parseImageConfigParams(c)

	// Assert results
	require.NoError(t, err)
	expected := &imageproxy.ImageConfig{
		Width:        300,
		Height:       200,
		ResizingType: imageproxy.ResizingTypeFit,
	}
	require.Equal(t, expected, result)
}

func Test_parseImageConfigParams_NoParamsSet_Success(t *testing.T) {
	// Setup test
	c, _ := setupGinTest("", "", nil, "")

	// Call function
	result, err := parseImageConfigParams(c)

	// Assert results
	require.NoError(t, err)
	require.Nil(t, result)
}
