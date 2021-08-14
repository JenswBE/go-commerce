package handler

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

func Test_parseImageConfigsParam_AllParamsSet_Success(t *testing.T) {
	// Setup test
	c, _ := setupGinTest(t, "", "?"+fixtures.ImageConfigQuery, nil, nil)

	// Call function
	result, err := parseImageConfigsParam(c)

	// Assert results
	require.NoError(t, err)
	require.Equal(t, fixtures.ImageConfigMap(), result)
}

func Test_parseImageConfigsParam_NoParamsSet_Success(t *testing.T) {
	// Setup test
	c, _ := setupGinTest(t, "", "", nil, nil)

	// Call function
	result, err := parseImageConfigsParam(c)

	// Assert results
	require.NoError(t, err)
	require.Nil(t, result)
}
