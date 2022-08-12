package admin

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/stretchr/testify/require"
)

func Test_parseFilesFromMultipart_Success(t *testing.T) {
	// Setup test
	body, writer := fixtures.MultipartMultipleFiles()
	req, _ := http.NewRequest("", "", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Call function
	files, err := parseFilesFromMultipart(req)

	// Assert results
	require.NoError(t, err)
	require.Equal(t, fixtures.MultipartMultipleFilesMap(), files)
}

func Test_parseFilesFromMultipart_Failure(t *testing.T) {
	// Setup test
	req, _ := http.NewRequest("", "", bytes.NewBufferString("invalid"))

	// Call function
	files, err := parseFilesFromMultipart(req)

	// Assert results
	require.Error(t, err)
	require.Nil(t, files)
}
