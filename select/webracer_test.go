package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	
	t.Run("check working urls", func(t *testing.T){
		slowServer := MakeServer(20 * time.Millisecond)
		fastServer := MakeServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("check timeout", func(t *testing.T) {
		slowServer := MakeServer(20 * time.Second)
		fastServer := MakeServer(15 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		got := Racer(slowServer.URL, fastServer.URL)

		if got != http.ErrHandlerTimeout.Error() {
			t.Fatalf("expected a timeout error and didn't get one")
		}
	})
	
}

func MakeServer(lag time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(lag)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}