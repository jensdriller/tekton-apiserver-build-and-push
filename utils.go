package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Backend struct {
	Count            int    `json="count"`
	LastRequestEpoch int64  `json="last_request"`
	Host             string `json="host"`
}

func fetchBackendHost(url string) (string, error) {
	req, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("request to backend service failed. Err: %s", err)
	}
	defer req.Body.Close()

	body, _ := ioutil.ReadAll(req.Body)

	var b Backend
	json.Unmarshal(body, &b)

	return b.Host, nil
}
