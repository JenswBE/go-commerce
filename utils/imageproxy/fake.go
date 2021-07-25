package imageproxy

import "fmt"

// Fake returns the input unchanged
type Fake struct{}

// NewFakeService creates a new fake image proxy service
func NewFakeService() *Fake {
	return &Fake{}
}

// GenerateURL appends the ImageConfig width to the sourceURL: <sourceURL>/<width>
func (fake *Fake) GenerateURL(sourceURL string, config ImageConfig) (string, error) {
	return fmt.Sprintf(`%s/%d`, sourceURL, config.Width), nil
}
