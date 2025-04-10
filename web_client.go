package eliotlibs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
		err = json.NewDecoder(resp.Body).Decode(&result)
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
		err = json.NewDecoder(resp.Body).Decode(&result)
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
