package imageproxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseImageConfig_Success(t *testing.T) {
	config, err := ParseImageConfig("300", "200", string(ResizingTypeFit))
	require.NoError(t, err)
	require.NotNil(t, config)
}

func Test_ParseImageConfig_InvalidWidth_Failure(t *testing.T) {
	config, err := ParseImageConfig("invalid", "200", string(ResizingTypeFill))
	require.Error(t, err)
	require.Nil(t, config)
}

func Test_ParseImageConfig_InvalidHeight_Failure(t *testing.T) {
	config, err := ParseImageConfig("300", "invalid", string(ResizingTypeFill))
	require.Error(t, err)
	require.Nil(t, config)
}

func Test_ParseImageConfig_InvalidResizingType_Failure(t *testing.T) {
	config, err := ParseImageConfig("300", "200", "invalid")
	require.Error(t, err)
	require.Nil(t, config)
}

func Test_ParseImageConfig_NegatifWidth_Failure(t *testing.T) {
	config, err := ParseImageConfig("-300", "200", string(ResizingTypeFill))
	require.Error(t, err)
	require.Nil(t, config)
}

func Test_ParseImageConfig_NegatifHeight_Failure(t *testing.T) {
	config, err := ParseImageConfig("300", "-200", string(ResizingTypeFill))
	require.Error(t, err)
	require.Nil(t, config)
}
