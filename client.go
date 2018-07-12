package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	"strconv"
)

type WavesNodeClient struct {
	Host   string
	Port   uint
	ApiKey string
	Method string
}

func (w *WavesNodeClient) DoRequest(uri string, method string) (*http.Response, error) {
	cl := http.Client{}

	url := fmt.Sprintf("http://%s:%s%s", w.Host, strconv.Itoa(int(w.Port)), uri)

	req, err := http.NewRequest(method, url, nil)

	req.Header.Set("api_key", w.ApiKey)

	if err != nil {
		return &http.Response{}, err
	}

	res, err := cl.Do(req)

	if err != nil {
		return res, err
	}

	return res, nil
}
