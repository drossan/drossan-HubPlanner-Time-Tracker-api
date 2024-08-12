package helpers

import (
	"io"
	"net/http"
)

func MakeHTTPRequest(method, url, apiToken, contentType string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", "Secuoyas Experiences - Time Tracking (daniel.rossello@secuoyas.com)")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
