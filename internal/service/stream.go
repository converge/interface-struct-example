package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Streamer struct {
	HttpClient HTTPClientProvider
}

func NewStreamer(httpClient HTTPClientProvider) Streamer {
	return Streamer{
		HttpClient:  httpClient,
	}
}


func (streamer Streamer) StartStream() error {

	out, _ := os.Create("output.txt")
	defer out.Close()

	r, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)

	resp, _ := streamer.HttpClient.Do(r)
	defer resp.Body.Close()

	// sendDataToStreamer

	n, _ := io.Copy(out, resp.Body)
	fmt.Println(n)

	return nil
}
