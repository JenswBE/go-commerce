package imageproxy_test

import (
	"testing"

	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/stretchr/testify/require"
)

func Test_ParseImageConfig_Success(t *testing.T) {
	config, err := imageproxy.ParseImageConfig("300", "200", string(imageproxy.ResizingTypeFit))
	require.NoError(t, err)
	require.NotNil(t, config)
}

func Test_ParseImageConfig_InvalidWidth_Failure(t *testing.T) {
	config, err := imageproxy.ParseImageConfig("invalid", "200", string(imageproxy.ResizingTypeFill))
	require.Error(t, err)
	require.Zero(t, config)
}

func Test_ParseImageConfig_InvalidHeight_Failure(t *testing.T) {
	config, err := imageproxy.ParseImageConfig("300", "invalid", string(imageproxy.ResizingTypeFill))
	require.Error(t, err)
	require.Zero(t, config)
}

func Test_ParseImageConfig_InvalidResizingType_Failure(t *testing.T) {
	config, err := imageproxy.ParseImageConfig("300", "200", "invalid")
	require.Error(t, err)
	require.Zero(t, config)
}

func Test_ParseImageConfig_NegatifWidth_Failure(t *testing.T) {
	config, err := imageproxy.ParseImageConfig("-300", "200", string(imageproxy.ResizingTypeFill))
	require.Error(t, err)
	require.Zero(t, config)
}

func Test_ParseImageConfig_NegatifHeight_Failure(t *testing.T) {
	config, err := imageproxy.ParseImageConfig("300", "-200", string(imageproxy.ResizingTypeFill))
	require.Error(t, err)
	require.Zero(t, config)
}
