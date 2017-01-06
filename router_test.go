package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterHealthCheck(t *testing.T) {
	router := NewNamegenRouter()
	resp := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/health", nil)

	router.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("Should have gotten 200 OK, but instead got %d", resp.Code)
	}
}

func TestRouterHealthCheckRedirectsWithCaps(t *testing.T) {
	router := NewNamegenRouter()
	resp := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/HEAltH", nil)

	router.ServeHTTP(resp, req)

	if resp.Code != 301 {
		t.Errorf("Expected to be redirected, but instead received status code %d", resp.Code)
	}
}
