package service

import (
	"net/http"
	"testing"
)

type MockHTTPClient struct {}

func (mock *MockHTTPClient) Do(_ *http.Request) (*http.Response, error) {
	return nil, nil
}

func TestNewDownloader(t *testing.T) {
	mockHTTP := MockHTTPClient{}
	s := NewDownloader(&mockHTTP)

	_ = s.GetIem()
}