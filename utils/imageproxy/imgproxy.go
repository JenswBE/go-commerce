package imageproxy

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// ImgProxy provides helper functions for working with imgproxy.
// See https://imgproxy.net/
type ImgProxy struct {
	allowedConfigs map[ImageConfig]bool
	baseURL        string
	key            []byte
	salt           []byte
}

// NewImgProxyService creates a new ImgProxy service
func NewImgProxyService(baseURL, keyHex, saltHex string, allowedConfigs []ImageConfig) (*ImgProxy, error) {
	// Restructure allowed configs for better lookup performance
	allowedConfigsMap := make(map[ImageConfig]bool, len(allowedConfigs))
	for _, config := range allowedConfigs {
		allowedConfigsMap[config] = true
	}

	// Validate base URL
	_, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %s", err.Error())
	}

	// Parse key
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex-encoded key: %s", err.Error())
	}

	// Parse salt
	salt, err := hex.DecodeString(saltHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex-encoded salt: %s", err.Error())
	}

	// Build service
	svc := &ImgProxy{
		allowedConfigs: allowedConfigsMap,
		baseURL:        baseURL,
		key:            key,
		salt:           salt,
	}
	return svc, nil
}

// GenerateURL generates a signed URL for imgproxy.
// See https://docs.imgproxy.net/generating_the_url
func (imgproxy *ImgProxy) GenerateURL(sourceURL string, config ImageConfig) (string, error) {
	// Check if config is allowed
	if len(imgproxy.allowedConfigs) > 0 && !imgproxy.allowedConfigs[config] {
		return "", errors.New("unsupported ImageConfig")
	}

	// Defaults
	gravity := "sm"
	enlarge := 1
	extension := "png"

	// Build path
	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(sourceURL))
	resizingType := strings.ToLower(string(config.ResizingType))
	path := fmt.Sprintf("/rt:%s/w:%d/h:%d/g:%s/el:%d/%s.%s",
		resizingType, config.Width, config.Height, gravity, enlarge, encodedURL, extension)

	// Calculate signature
	mac := hmac.New(sha256.New, imgproxy.key)
	mac.Write(imgproxy.salt)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	// Build result
	return imgproxy.baseURL + signature + path, nil
}
