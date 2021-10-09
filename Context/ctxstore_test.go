package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (store *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return store.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("gets the response", func(t *testing.T) {
		data := "Hello"
		store := &SpyStore{response: data}
		srv := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		srv.ServeHTTP(response, request)

		if store.cancelled {
			t.Error("store was cancelled")
		}

		if response.Body.String() != "Hello" {
			t.Errorf("got %s want %s", response.Body.String(), data)
		}

	})

	t.Run("tells store to cancel if the request is cancelled", func(t *testing.T) {
		data := "Hello"
		store := &SpyStore{response: data}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(ctx)

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}
