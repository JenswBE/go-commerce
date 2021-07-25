package imageproxy

//Service interface
type Service interface {
	// GenerateURL generates a signed URL for the image proxy
	GenerateURL(sourceURL string, config ImageConfig) (string, error)
}
