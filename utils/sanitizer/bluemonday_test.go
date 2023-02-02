package sanitizer_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/JenswBE/go-commerce/utils/sanitizer"
)

func Test_Sanitize(t *testing.T) {
	service := sanitizer.NewBluemondayService()
	tests := []struct {
		given     string
		expected  string
		sanitizer func(string) string
	}{
		{
			given:     "test",
			expected:  "test",
			sanitizer: service.String,
		},
		{
			given:     "<p>test</p>",
			expected:  "test",
			sanitizer: service.String,
		},
		{
			given:     "test",
			expected:  "test",
			sanitizer: service.ContentHTML,
		},
		{
			given:     "<p><strong>test<br/>test2</strong></p>",
			expected:  "<p><strong>test<br/>test2</strong></p>",
			sanitizer: service.ContentHTML,
		},
		{
			given:     "<p><strong>test<br/><script>DoSomething()</script></strong></p>",
			expected:  "<p><strong>test<br/></strong></p>",
			sanitizer: service.ContentHTML,
		},
	}

	for i, subtest := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			require.Equal(t, subtest.expected, subtest.sanitizer(subtest.given))
		})
	}
}
