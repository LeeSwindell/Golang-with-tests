package poker

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGETPlayers(t *testing.T) {

	playerStore := StubPlayerStore{
		map[string]int{"jo": 20, "billy":50},
		[]string{},
		League{},
	}
	server := NewPlayerServer(&playerStore)

	for player, score  := range playerStore.Scores {
		t.Run(player, func(t *testing.T) {
			request := NewGetScoreRequest(player)
			response := httptest.NewRecorder()
	
			server.ServeHTTP(response, request)
	
			AssertResponseBody(t, response.Body.String(), fmt.Sprint(score))
			AssertStatus(t, response.Code, http.StatusOK)
		})
	}

	t.Run("404 on missing players", func(t *testing.T) {
		request := NewGetScoreRequest("brian")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		AssertStatus(t, got, want)
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

		AssertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("update existing player", func(t *testing.T) {
		request := NewPostWinRequest("jo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		
		AssertStatus(t, response.Code, http.StatusAccepted)
		got := playerStore.WinCalls[1]
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

		request := NewLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)
		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, *response, jsonContentType)
	})
}
