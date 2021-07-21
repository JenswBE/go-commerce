package imageproxy

//Service interface
type Service interface {
	// Encode converts an UUID to a short ID
	GenerateURL(sourceURL string, config ImageConfig) (string, error)
}
