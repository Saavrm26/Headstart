package Utils

import (
	"bytes"
	"mime/multipart"
)

func MultipartFileBuffer(fileHeader *multipart.FileHeader) (*bytes.Buffer, error) {
	fileReader, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	fileBytes := make([]byte, fileHeader.Size)
	_, err = fileReader.Read(fileBytes)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(fileBytes)
	return buf, nil
}
