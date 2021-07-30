package imageproxy

import "fmt"

type ResizingType string

const ResizingTypeFit ResizingType = "fit"
const ResizingTypeFill ResizingType = "fill"

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