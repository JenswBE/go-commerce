package imageproxy

import "fmt"

type ResizingType string

const (
	ResizingTypeFit  ResizingType = "FIT"
	ResizingTypeFill ResizingType = "FILL"
)

func ParseResizingType(input string) (ResizingType, error) {
	switch ResizingType(input) {
	case ResizingTypeFit:
		return ResizingTypeFit, nil
	case ResizingTypeFill:
		return ResizingTypeFill, nil
	default:
		return ResizingType(""), fmt.Errorf(`invalid ResizingType: %s`, input)
	}
}
