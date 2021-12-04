package sanitizer

// Fake returns the input unchanged
type Fake struct{}

// NewFakeService creates a new fake sanitizer service
func NewFakeService() *Fake {
	return &Fake{}
}

// Sanitizes input to a plain string
func (f *Fake) String(input string) string {
	return input
}

// Sanitizes input to a very restricted HTML subset for content of type HTML
func (f *Fake) ContentHTML(input string) string {
	return input
}
