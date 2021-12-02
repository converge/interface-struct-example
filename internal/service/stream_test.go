package service

import "testing"

func TestNewStreamer(t *testing.T) {
	mockHTTP := MockHTTPClient{}
	s := NewStreamer(&mockHTTP)

	_ = s.StartStream()
}