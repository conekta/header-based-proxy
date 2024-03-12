package header_based_proxy_test

import (
	"net/http"
	"testing"
)

func TestProxy(t *testing.T) {
	
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
