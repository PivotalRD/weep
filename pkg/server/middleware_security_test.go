package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBrowserFilterMiddleware_NoHostHeader(t *testing.T) {
	// Create a request with no Host header
	req := httptest.NewRequest("GET", "http://127.0.0.1:9091/", nil)
	req.Host = ""

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Create a handler to be protected by the middleware
	protectedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create the middleware
	middleware := BrowserFilterMiddleware(protectedHandler)

	// Serve the request
	middleware.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
