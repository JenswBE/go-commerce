package sanitizer

//Service interface
type Service interface {
	// Sanitizes input to a plain string
	String(input string) string

	// Sanitizes input to a very restricted HTML subset for content of type HTML
	ContentHTML(input string) string
}
