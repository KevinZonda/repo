package utils

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
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

func HttpGetYaml[T any](url string) (t T, err error) {
	hc := http.DefaultClient
	resp, err := hc.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = yaml.NewDecoder(resp.Body).Decode(&t)
	return
}

func HttpGetStr(url string) (t string, err error) {
	hc := http.DefaultClient
	resp, err := hc.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err == nil {
		t = string(bs)
	}
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
