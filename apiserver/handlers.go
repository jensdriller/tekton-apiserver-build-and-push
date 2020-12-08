package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

var counter = Counter{
	Count:       1,
	LastRequest: time.Now().Unix(),
}

type Counter struct {
	Count       int    `json:"count"`
	LastRequest int64  `json:"last_request"`
	Host        string `json:"host"`
	Backend     string `json:"backend"`
}

func getCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	counter.Count++
	counter.LastRequest = time.Now().Unix()
	counter.Host, _ = os.Hostname()
	var err error
	// counter.Backend, err = fetchBackendHost("http://backend")
	if err != nil {
		panic(err) // probably want to log fatal rather than crashing the server
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(counter); err != nil {
		panic(err)
	}
}
