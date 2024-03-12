package header_based_proxy_test

import (
	"context"
	proxy "github.com/conekta/header-based-proxy"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProxy(t *testing.T) {
	cfg := proxy.CreateConfig()
	cfg.Header = "X-Femsa-Migrated"
	cfg.Mapping["true"] = "https://api.femsa.io"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := proxy.New(ctx, next, cfg, "proxy-test")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	print(req)
}
