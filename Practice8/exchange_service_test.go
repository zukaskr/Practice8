package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRate(t *testing.T) {
	t.Run("Success Scenario", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"rate":0.92}`))
		}))
		defer server.Close()

		svc := NewExchangeService(server.URL)
		rate, err := svc.GetRate("USD", "EUR")
		assert.NoError(t, err)
		assert.Equal(t, 0.92, rate)
	})

	t.Run("API Error 404", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"not found"}`))
		}))
		defer server.Close()

		svc := NewExchangeService(server.URL)
		_, err := svc.GetRate("INVALID", "PHP")
		assert.Error(t, err)
	})
}
