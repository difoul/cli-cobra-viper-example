package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func httpRequest(method string, url string, authorization string, body io.Reader, contentType string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%s %s returned %d status", method, url, resp.StatusCode))
	}
	return bodyr, nil
}

func HttpGet(url string, username string, password string) ([]byte, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return httpRequest(http.MethodGet, url, fmt.Sprintf("Basic %s", auth), nil, "application/json")
}

func HttpDelete(url string, username string, password string) ([]byte, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return httpRequest(http.MethodDelete, url, fmt.Sprintf("Basic %s", auth), nil, "application/json")
}

func HttpPost(url string, username string, password string, body string) ([]byte, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return httpRequest(http.MethodPost, url, fmt.Sprintf("Basic %s", auth), bytes.NewBuffer([]byte(body)), "application/json")
}

func HttpPut(url string, username string, password string, body string) ([]byte, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return httpRequest(http.MethodPut, url, fmt.Sprintf("Basic %s", auth), bytes.NewBuffer([]byte(body)), "application/json")
}
