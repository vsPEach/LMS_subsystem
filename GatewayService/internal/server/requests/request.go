package requests

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func Request(body []byte, url string) error {
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "fuck me PLS")
	client := &http.Client{Timeout: time.Minute}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return err
}
