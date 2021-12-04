package sanitizer

import "github.com/microcosm-cc/bluemonday"

// Bluemonday uses the bluemonday package for sanitizing input
type Bluemonday struct {
	stringPolicy      *bluemonday.Policy
	contentHTMLPolicy *bluemonday.Policy
}

// NewBluemondayService creates a new bluemonday sanitizer service
func NewBluemondayService() *Bluemonday {
	contentHTMLPolicy := bluemonday.NewPolicy()
	contentHTMLPolicy.AllowElements(
		// Basics
		"p",
		"br",

		// Formatting
		"strong", // Bold
		"em",     // Italics
		"u",      // Underlined

		// Lists
		"ol",
		"ul",
		"li",
	)

	return &Bluemonday{
		stringPolicy:      bluemonday.StrictPolicy(),
		contentHTMLPolicy: contentHTMLPolicy,
	}
}

// Sanitizes input to a plain string
func (b *Bluemonday) String(input string) string {
	return b.stringPolicy.Sanitize(input)
}

// Sanitizes input to a very restricted HTML subset for content of type HTML
func (b *Bluemonday) ContentHTML(input string) string {
	return b.contentHTMLPolicy.Sanitize(input)
}
