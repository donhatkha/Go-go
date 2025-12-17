package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		envValue     string
		expected     string
	}{
		{
			name:         "Environment variable exists",
			key:          "TEST_VAR",
			defaultValue: "default",
			envValue:     "custom",
			expected:     "custom",
		},
		{
			name:         "Environment variable does not exist",
			key:          "NON_EXISTENT_VAR",
			defaultValue: "default",
			envValue:     "",
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				t.Setenv(tt.key, tt.envValue)
			}

			result := getEnv(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("getEnv(%s, %s) = %s; want %s", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

func TestCORSMiddleware(t *testing.T) {
	handler := corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	// Check CORS headers
	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("CORS Allow-Origin header not set correctly")
	}

	if w.Header().Get("Access-Control-Allow-Methods") != "GET, POST, PUT, DELETE, OPTIONS" {
		t.Error("CORS Allow-Methods header not set correctly")
	}

	if w.Header().Get("Access-Control-Allow-Headers") != "Content-Type, Authorization" {
		t.Error("CORS Allow-Headers header not set correctly")
	}
}

func TestCORSMiddlewareOptions(t *testing.T) {
	handler := corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	req := httptest.NewRequest("OPTIONS", "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	// OPTIONS should return 200 without executing the handler
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200 for OPTIONS, got %d", w.Code)
	}

	// Body should be empty for OPTIONS request
	if w.Body.String() != "" {
		t.Error("Expected empty body for OPTIONS request")
	}
}

func TestRespondJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := Response{
		Success: true,
		Message: "Test message",
		Data:    "Test data",
	}

	respondJSON(w, http.StatusOK, data)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}

func TestRespondError(t *testing.T) {
	w := httptest.NewRecorder()
	message := "Test error message"

	respondError(w, http.StatusBadRequest, message)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}
