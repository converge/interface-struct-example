package service

import (
	"example-delete-it/internal/mocks"
	"testing"
)

func TestNewStreamer(t *testing.T) {
	mockHTTP := mocks.MockHTTPClient{}
	s := NewStreamer(&mockHTTP)

	_ = s.StartStream()
}