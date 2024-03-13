package header_based_proxy_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	proxy "github.com/conekta/header-based-proxy"
	"github.com/stretchr/testify/assert"
)

func TestProxy(t *testing.T) {
	tests := []struct {
		name        string
		headerName  string
		headerValue string
		expected    string
	}{
		{
			name:        "Proxy enabled",
			headerName:  "X-Femsa-Migrated",
			headerValue: "true",
			expected:    "api.digitalfemsa.io",
		},
		{
			name:        "Proxy disabled",
			headerValue: "false",
			headerName:  "X-Femsa-Migrated",
			expected:    "api.conekta.io",
		},
		{
			name:        "Header Empty",
			headerName:  "X-Femsa-Migrated",
			headerValue: "",
			expected:    "api.conekta.io",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := proxy.CreateConfig()
			cfg.Header = "X-Femsa-Migrated"
			cfg.Mapping["true"] = "https://api.digitalfemsa.io"

			ctx := context.Background()
			next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

			handler, err := proxy.New(ctx, next, cfg, "proxy-test")
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.conekta.io", nil)

			req.Header.Add(tt.headerName, tt.headerValue)

			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(recorder, req)

			assert.EqualValues(t, tt.expected, req.Host)
		})
	}
}
