package tools

import (
	"fmt"
	"io"
	"net/http"
)

func ShortenURL(longURL string) (string, error) {
	url := fmt.Sprintf("http://tinyurl.com/api-create.php?url=%s", longURL)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
