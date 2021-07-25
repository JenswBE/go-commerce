package imageproxy_test

import (
	"testing"

	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/stretchr/testify/require"
)

func Test_NewImgProxyService_Success(t *testing.T) {
	service, err := imageproxy.NewImgProxyService("http://localhost", "736563726574", "68656C6C6F", nil)
	require.NoError(t, err)
	require.NotNil(t, service)
}

func Test_NewImgProxyService_InvalidBaseURL_Failure(t *testing.T) {
	service, err := imageproxy.NewImgProxyService("\n", "736563726574", "68656C6C6F", nil)
	require.Error(t, err)
	require.Nil(t, service)
}

func Test_NewImgProxyService_InvalidKey_Failure(t *testing.T) {
	service, err := imageproxy.NewImgProxyService("http://localhost", "invalid", "68656C6C6F", nil)
	require.Error(t, err)
	require.Nil(t, service)
}

func Test_NewImgProxyService_InvalidSalt_Failure(t *testing.T) {
	service, err := imageproxy.NewImgProxyService("http://localhost", "736563726574", "invalid", nil)
	require.Error(t, err)
	require.Nil(t, service)
}

func Test_GenerateURL_Success(t *testing.T) {
	// Test based on example at https://docs.imageproxy.net/signing_the_url

	// Create service
	key := "736563726574" // secret
	salt := "68656C6C6F"  // hello
	config := imageproxy.ImageConfig{Width: 300, Height: 400, ResizingType: imageproxy.ResizingTypeFill}
	service, err := imageproxy.NewImgProxyService("https://imageproxy.test/", key, salt, []imageproxy.ImageConfig{config})
	require.NoError(t, err)

	// Generate URL
	imgURL, err := service.GenerateURL("http://example.com/images/curiosity.jpg", config)

	// Assert results
	require.NoError(t, err)
	expected := "https://imageproxy.test/UTshYhaXeFJ518RnptmRB3KkBbuxxukwfPUXpnH2BBU/fill/300/400/sm/1/aHR0cDovL2V4YW1wbGUuY29tL2ltYWdlcy9jdXJpb3NpdHkuanBn.png"
	require.Equal(t, expected, imgURL)
}

func Test_GenerateURL_UnsupportedImageConfig_Success(t *testing.T) {
	// Create service
	config := imageproxy.ImageConfig{Width: 300, Height: 400, ResizingType: imageproxy.ResizingTypeFill}
	service, err := imageproxy.NewImgProxyService("", "", "", []imageproxy.ImageConfig{config})
	require.NoError(t, err)

	// Generate URL
	config.Width = 512
	imgURL, err := service.GenerateURL("", config)

	// Assert results
	require.Error(t, err)
	require.Empty(t, imgURL)
}
