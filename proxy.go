package header_based_proxy

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

// Config defines the plugin configuration.
type Config struct {
	Header  string            `json:"header,omitempty"`  // target header
	Mapping map[string]string `json:"mapping,omitempty"` // mapping pairs (regex, target service)
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Header:  "",
		Mapping: make(map[string]string),
	}
}

// CustomProxy defines the proxy struct
type CustomProxy struct {
	config *Config
	next   http.Handler
	name   string
}

// New creates a new CustomProxy plugin instance
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Header) == 0 {
		return nil, fmt.Errorf("header cannot be empty")
	}

	if len(config.Mapping) == 0 {
		return nil, fmt.Errorf("mapping cannot be empty")
	}

	return &CustomProxy{
		config: config,
		next:   next,
		name:   name,
	}, nil
}

// Process the requests to verify if they match with defined mapping patterns
func (a *CustomProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for pattern, destination := range a.config.Mapping {
		matched, _ := regexp.MatchString(pattern, req.Header.Get(a.config.Header))

		if matched {
			destinationUrl, err := url.Parse(destination)

			if err != nil {
				continue
			}

			proxy := httputil.NewSingleHostReverseProxy(destinationUrl)
			req.URL = destinationUrl
			req.Host = destinationUrl.Host
			proxy.ServeHTTP(rw, req)
			return
		}
	}
	a.next.ServeHTTP(rw, req)
}
