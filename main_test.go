package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	} else {
		t.Logf("got 200")
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	} else {
		t.Logf("got 200")
	}
}

func TestPrimeHandler(t *testing.T) {
	tests := []struct {
		num      string
		expected string
		status   int
	}{
		{"17", "17 is Prime", http.StatusOK},
		{"21", "21 is NOT Prime", http.StatusOK},
		{"1", "1 is NOT Prime", http.StatusOK},
		{"-17", "-17 is NOT Prime", http.StatusOK},
		{"0", "0 is NOT Prime", http.StatusOK},
		{"string", "num parameter must be an integer", http.StatusBadRequest},
		{"17a", "num parameter must be an integer", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/prime/?num="+tt.num, nil)
			rec := httptest.NewRecorder()
			primeHandler(rec, req)
			if rec.Code != tt.status {
				t.Errorf("expected 200, got %d", rec.Code)
			} else {
				t.Logf("got 200")
			}

			if !strings.Contains(rec.Body.String(), tt.expected) {
				t.Errorf("expected %s, got %s", tt.expected, rec.Body.String())
			} else {
				t.Logf("got 200")
			}
		})
	}
}
