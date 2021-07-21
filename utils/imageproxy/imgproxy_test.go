package imageproxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenerateURL_Success(t *testing.T) {
	// Test based on example at https://docs.imgproxy.net/signing_the_url

	// Create service
	key := "736563726574" // secret
	salt := "68656C6C6F"  // hello
	config := ImageConfig{Width: 300, Height: 400, ResizingType: ResizingTypeFill}
	service, err := NewImgProxyService("https://imgproxy.test/", key, salt, []ImageConfig{config})
	require.NoError(t, err)

	// Generate URL
	imgURL, err := service.GenerateURL("http://example.com/images/curiosity.jpg", config)

	// Assert results
	require.NoError(t, err)
	expected := "https://imgproxy.test/UTshYhaXeFJ518RnptmRB3KkBbuxxukwfPUXpnH2BBU/fill/300/400/sm/1/aHR0cDovL2V4YW1wbGUuY29tL2ltYWdlcy9jdXJpb3NpdHkuanBn.png"
	require.Equal(t, expected, imgURL)
}

func Test_GenerateURL_UnsupportedImageConfig_Success(t *testing.T) {
	// Create service
	config := ImageConfig{Width: 300, Height: 400, ResizingType: ResizingTypeFill}
	service, err := NewImgProxyService("", "", "", []ImageConfig{config})
	require.NoError(t, err)

	// Generate URL
	config.Width = 512
	imgURL, err := service.GenerateURL("", config)

	// Assert results
	require.Error(t, err)
	require.Empty(t, imgURL)
}
