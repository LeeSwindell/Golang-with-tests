package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)


type StubPlayerStore struct{
	scores map[string]int
	winCalls []string
	league League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func(s *StubPlayerStore) GetLeague() League {
	return s.league
}

func TestGETPlayers(t *testing.T) {

	playerStore := StubPlayerStore{
		map[string]int{"jo": 20, "billy":50},
		[]string{},
		League{},
	}
	server := NewPlayerServer(&playerStore)

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
		League{},
	}
	server := NewPlayerServer(&playerStore)

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
	
	t.Run("return league table as json", func(t *testing.T) {
		
		wantedLeague := League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, *response, jsonContentType)
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

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league League) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	
	if err != nil {
		t.Fatalf("couldn't parse %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func assertContentType(t testing.TB, response httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %q, got %v", want, response.Result().Header)
	}
}
