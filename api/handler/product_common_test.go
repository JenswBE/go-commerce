package handler

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/JenswBE/go-commerce/utils/imageproxy"
	"github.com/stretchr/testify/require"
)

func Test_parseFilesFromMultipart_Success(t *testing.T) {
	// Setup test
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	test1, _ := writer.CreateFormFile("file", "test1.jpg")
	test1.Write([]byte("test-1-jpg"))
	test2, _ := writer.CreateFormFile("file", "test2.png")
	test2.Write([]byte("test-2-png"))
	writer.Close()
	req, _ := http.NewRequest("", "", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Call function
	files, err := parseFilesFromMultipart(req)

	// Assert results
	require.NoError(t, err)
	expected := map[string][]byte{
		"test1.jpg": []byte("test-1-jpg"),
		"test2.png": []byte("test-2-png"),
	}
	require.Equal(t, expected, files)
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

func Test_parseImageConfigParams_AllParamsSet_Success(t *testing.T) {
	// Setup test
	c, _ := setupGinTest(t, "", "?"+fixtures.ImageConfigQuery, nil, nil)

	// Call function
	result, err := parseImageConfigParams(c)

	// Assert results
	require.NoError(t, err)
	expected := &imageproxy.ImageConfig{
		Width:        300,
		Height:       200,
		ResizingType: imageproxy.ResizingTypeFill,
	}
	require.Equal(t, expected, result)
}

func Test_parseImageConfigParams_NoParamsSet_Success(t *testing.T) {
	// Setup test
	c, _ := setupGinTest(t, "", "", nil, nil)

	// Call function
	result, err := parseImageConfigParams(c)

	// Assert results
	require.NoError(t, err)
	require.Nil(t, result)
}
