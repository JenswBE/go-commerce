package main

import (
	"testing"

	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/stretchr/testify/require"
)

func Test_parseAllowedImageConfigs_MultipleConfigs_Success(t *testing.T) {
	// Call helper
	input := "100:100:FIT,300:200:FIT,512:512:FILL"
	result, err := parseAllowedImageConfigs(input)

	// Assert results
	require.NoError(t, err)
	expected := []imageproxy.ImageConfig{
		{
			Width:        100,
			Height:       100,
			ResizingType: imageproxy.ResizingTypeFit,
		},
		{
			Width:        300,
			Height:       200,
			ResizingType: imageproxy.ResizingTypeFit,
		},
		{
			Width:        512,
			Height:       512,
			ResizingType: imageproxy.ResizingTypeFill,
		},
	}
	require.Equal(t, expected, result)
}

func Test_parseAllowedImageConfigs_Wildcard_Success(t *testing.T) {
	// Call helper
	result, err := parseAllowedImageConfigs("*")

	// Assert results
	require.NoError(t, err)
	require.Len(t, result, 0)
}
