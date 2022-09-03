package main

import (
	"demo/http-server/api"
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
