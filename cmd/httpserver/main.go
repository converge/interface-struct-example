package main

import (
	"example-delete-it/internal/service"
	"net/http"
)

func main() {

	httpClient := http.DefaultClient

	service.NewDownloader(httpClient)

	service.NewStreamer(httpClient)

}
