package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)


type StubPlayerStore struct{
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {

	playerStore := StubPlayerStore{
		map[string]int{"jo": 20, "billy":50},
		[]string{},
	}
	server := &PlayerServer{&playerStore}

	for player, score  := range playerStore.scores {
		t.Run(player, func(t *testing.T) {
			request := newGetScoreRequest(player)
			response := httptest.NewRecorder()
	
			server.ServeHTTP(response, request)
	
			assertResponseBody(t, response.Body.String(), fmt.Sprint(score))
			assertStatus(t, response.Code, http.StatusOK)
		})
	}

	t.Run("404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("brian")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		assertStatus(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {

	playerStore := StubPlayerStore{
		map[string]int{"jo": 20, "billy":50},
		[]string{},
	}
	server := &PlayerServer{&playerStore}

	t.Run("check new player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/brian", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("update existing player", func(t *testing.T) {
		request := newPostWinRequest("jo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Code, http.StatusAccepted)
		got := playerStore.winCalls[1]
		want := "jo"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := &PlayerServer{&store}

	t.Run("return status ok", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func newGetScoreRequest(name string) (*http.Request) {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}

func newPostWinRequest(name string) (*http.Request) {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got response %q, wanted response %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, wanted %d", got, want)
	}
}