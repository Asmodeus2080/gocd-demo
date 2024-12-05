package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	// Use the handler from main.go
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Call the actual handler
	MainHandler(rec, req)

	// Check the status code
	if rec.Code != http.StatusAccepted {
		t.Errorf("expected status %d, got %d", http.StatusAccepted, rec.Code)
	}

	// Check the response body
	expectedBody := `"hello this is a very dummy project"` + "\n" // JSON encoding adds a newline
	if rec.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rec.Body.String())
	}

	// Check the content type
	expectedContentType := "application/json"
	if rec.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("expected content type %q, got %q", expectedContentType, rec.Header().Get("Content-Type"))
	}
}
