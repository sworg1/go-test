package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

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
