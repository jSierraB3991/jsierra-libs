package eliotlibs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type HeaderRequest struct {
	Key   string
	Value string
}

func Get(urlBase, uri string, result interface{}, headers []HeaderRequest) error {

	req, err := http.NewRequest("GET", urlBase+uri, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&result)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New(resp.Status)
}

func Post(urlBase, uri string, jsonData []byte, result interface{}, headers []HeaderRequest) error {
	var bodyIoReader io.Reader
	if jsonData != nil {
		bodyIoReader = bytes.NewBuffer(jsonData)
	}
	req, err := http.NewRequest("POST", urlBase+uri, bodyIoReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&result)
		if err != nil {
			return err
		}
		return nil
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error al leer la respuesta:", err)
			return err
		}

		// Imprimir el cuerpo de la respuesta
		fmt.Println("Respuesta:", string(body))
	}
	return errors.New(resp.Status)
}

func PostFile(urlBase, uri string, fileHeader *multipart.FileHeader, result interface{}, headers []HeaderRequest) error {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffer to hold the multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field and copy the file content into it
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Close the writer to finalize the multipart form data
	err = writer.Close()
	if err != nil {
		return err
	}

	// Create the POST request
	req, err := http.NewRequest("POST", urlBase+uri, body)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&result)
		if err != nil {
			return err
		}
		return nil
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}
	fmt.Println("Response:", string(bodyResp))
	return errors.New(resp.Status)
}
