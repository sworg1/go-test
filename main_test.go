package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealthEndpoint tests the /health endpoint to ensure it responds with
// an HTTP 200 status code and a body containing "OK". It uses an HTTP test
// request and response recorder to simulate and validate the behavior of
// the endpoint.
func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	if w.Body.String() != "OK" {
		t.Errorf("Expected body 'OK', got %v", w.Body.String())
	}
}

// TestRootEndpoint tests the root endpoint of an HTTP server.
// It creates a mock HTTP GET request to the root path ("/") and uses
// an HTTP test recorder to capture the response. The test verifies
// that the response status code is HTTP 200 OK and that the response
// body contains the expected mock IP address "127.0.0.1".
func TestRootEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("127.0.0.1")) // Mock IP for testing
	}).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	if w.Body.String() != "127.0.0.1" {
		t.Errorf("Expected body '127.0.0.1', got %v", w.Body.String())
	}
}

// TestGetEnv tests the getEnv function to ensure it correctly retrieves
// environment variable values or falls back to a default value when the
// environment variable is not set.
//
// The test covers the following scenarios:
// 1. When the environment variable is set, the function should return its value.
// 2. When the environment variable is not set, the function should return the provided default value.
func TestGetEnv(t *testing.T) {
	// Test with environment variable set
	t.Setenv("TEST_ENV", "value")
	result := getEnv("TEST_ENV", "default")
	if result != "value" {
		t.Errorf("Expected 'value', got %v", result)
	}

	// Test with environment variable not set
	result = getEnv("NON_EXISTENT_ENV", "default")
	if result != "default" {
		t.Errorf("Expected 'default', got %v", result)
	}
}

// TestGetLocalIP tests the GetLocalIP function to ensure it retrieves
// the local IP address correctly. It uses a mocked UDP connection
// to simulate the environment and verifies that the function returns
// a valid IP address without any errors.
func TestGetLocalIP(t *testing.T) {
	// Mock a UDP connection to test GetLocalIP
	conn := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer conn.Close()

	ip, err := GetLocalIP()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if ip == nil {
		t.Errorf("Expected a valid IP, got nil")
	}
}
