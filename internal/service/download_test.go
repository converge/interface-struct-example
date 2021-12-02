package service

import (
	"example-delete-it/internal/mocks"
	"testing"
)


func TestNewDownloader(t *testing.T) {
	mockHTTP := mocks.MockHTTPClient{}
	s := NewDownloader(&mockHTTP)

	_ = s.GetIem()
}