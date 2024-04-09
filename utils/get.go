package utils

import (
	"encoding/json"
	"net/http"
)

func HttpGetJson[T any](url string) (t T, err error) {
	hc := http.DefaultClient
	resp, err := hc.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&t)
	return
}

func HttpGetUnmarshal(url string, v any) (err error) {
	hc := http.DefaultClient
	resp, err := hc.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return
}
