package mocks

import "net/http"

type MockHTTPClient struct {}

func (mock *MockHTTPClient) Do(_ *http.Request) (*http.Response, error) {
	return nil, nil
}
