package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("unable to fetch html: %v", err)
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("response contains error code: %v", resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "text/html" {
		return "", fmt.Errorf("incorrect content type: %v", resp.Header.Get("Content-Type"))
	}

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read html: %v", err)
	}
	return string(html), nil
}
