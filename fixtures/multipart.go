package fixtures

import (
	"bytes"
	"io"
	"mime/multipart"
)

func MultipartSingleFile() (*bytes.Buffer, *multipart.Writer) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	test, _ := writer.CreateFormFile("file", "test.jpg")
	_, _ = io.WriteString(test, "test-jpg") // We don't care for errors in fixtures
	writer.Close()
	return body, writer
}

func MultipartSingleFileMap() map[string][]byte {
	return map[string][]byte{"test.jpg": []byte("test-jpg")}
}

func MultipartMultipleFiles() (*bytes.Buffer, *multipart.Writer) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	test1, _ := writer.CreateFormFile("file", "test1.jpg")
	_, _ = io.WriteString(test1, "test-1-jpg") // We don't care for errors in fixtures
	test2, _ := writer.CreateFormFile("file", "test2.png")
	_, _ = io.WriteString(test2, "test-2-png") // We don't care for errors in fixtures
	writer.Close()
	return body, writer
}

func MultipartMultipleFilesMap() map[string][]byte {
	return map[string][]byte{
		"test1.jpg": []byte("test-1-jpg"),
		"test2.png": []byte("test-2-png"),
	}
}
