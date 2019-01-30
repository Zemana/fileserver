package main

import (
	"net/http"
	"endpoints"
	"time"
	"math/rand"
)

var routerMap = map[string] func(w http.ResponseWriter, r *http.Request) {
	"/sample/up/": endpoints.Up,
	"/sample/upload/": endpoints.Upload,
	"/sample/exists/": endpoints.Exists,
	"/sample/download/": endpoints.Download,
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for pat, fun := range routerMap {
		http.HandleFunc(pat, fun)
	}
	http.ListenAndServe(":8080", nil)
}
