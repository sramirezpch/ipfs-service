package service

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/sramirezpch/ipfs-writer/config"
)

type ImageService struct {
	baseUrl string
}

func NewImageService(c *config.Config) *ImageService {
	return &ImageService{baseUrl: c.ImageServiceUrl}
}

func (s *ImageService) SaveImage(file multipart.File, header *multipart.FileHeader) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", header.Filename)
	if err != nil {
		log.Fatalf("Could not create form file: %s", err)
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatalf("Could not copy the file to the reader: %s", err)
		return err
	}

	err = writer.Close()
	if err != nil {
		log.Fatalf("Could not finish writing the multipart request: %s", err)
		return err
	}

	req, reqErr := http.NewRequest("POST", fmt.Sprintf("%s/save", s.baseUrl), body)
	if reqErr != nil {
		log.Fatalf("Could not create the request: %s", reqErr)
		return reqErr
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Something happened while making the request: %s", err)
		return err
	}

	defer res.Body.Close()

	return nil
}
