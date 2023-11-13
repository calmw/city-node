package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func HttpGet(uri string, header map[string]string, data map[string]string) ([]byte, error) {

	if data != nil && len(data) > 0 {
		params := url.Values{}
		for k, v := range data {
			params.Set(k, fmt.Sprintf("%v", v))
		}
		uri = uri + "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", uri, nil)

	req.Header.Add("content-type", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}

func HttpGeLocation(url string) (error, []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, []byte{}
	}
	req.Header.Set("x-into-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsInVzZXJfbmFtZSI6InRlc3QifQ.riZtv7y-kexYAb7mXp6cpf9G-Flb1rb-2POtNQXXe8E")
	resp, err := client.Do(req)
	if err != nil {
		return err, []byte{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, []byte{}
	}

	return nil, body
}

func HttpPost(uri string, header map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}
	//req.Header.Add("TRON-PRO-API-KEY", tron.ApiKey)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}
