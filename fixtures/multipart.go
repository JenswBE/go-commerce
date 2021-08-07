package fixtures

import (
	"bytes"
	"mime/multipart"
)

func MultipartSingleFile() (*bytes.Buffer, *multipart.Writer) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	test, _ := writer.CreateFormFile("file", "test.jpg")
	test.Write([]byte("test-jpg"))
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
	test1.Write([]byte("test-1-jpg"))
	test2, _ := writer.CreateFormFile("file", "test2.png")
	test2.Write([]byte("test-2-png"))
	writer.Close()
	return body, writer
}

func MultipartMultipleFilesMap() map[string][]byte {
	return map[string][]byte{
		"test1.jpg": []byte("test-1-jpg"),
		"test2.png": []byte("test-2-png"),
	}
}
