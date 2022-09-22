package poker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

type StubPlayerStore struct{
	Scores map[string]int
	WinCalls []string
	League League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func(s *StubPlayerStore) GetLeague() League {
	return s.League
}

func NewGetScoreRequest(name string) (*http.Request) {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}

func NewPostWinRequest(name string) (*http.Request) {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return request
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got response %q, wanted response %q", got, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, wanted %d", got, want)
	}
}

func NewLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func GetLeagueFromResponse(t testing.TB, body io.Reader) (league League) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	
	if err != nil {
		t.Fatalf("couldn't parse %q into slice of Player, '%v'", body, err)
	}

	return
}

func AssertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func AssertContentType(t testing.TB, response httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %q, got %v", want, response.Result().Header)
	}
}

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("couldnt create temp file, err %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	
	return tmpfile, removeFile
}

func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func AssertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	amountGot := got.Amount
	if amountGot != want.Amount {
		t.Errorf("got amount %d, want %d", amountGot, want.Amount)
	}

	gotScheduledTime := got.At
	if gotScheduledTime != want.At {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, want.At)
	}
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

func AssertPlayerWins(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatal("expected a win and didnt get one")
	}

	if store.WinCalls[0] != winner {
		t.Errorf("got winner %q, wanted %q", store.WinCalls[0], winner)
	}
}

type GameSpy struct{
	StartedWith int
	FinishedWith string
	StartCalled bool
	FinishCalled bool
}

func(g *GameSpy) Start(numPlayers int) {
	g.StartedWith = numPlayers
	g.StartCalled = true
}

func(g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}

func AssertCorrectGame(t testing.TB, got *GameSpy, numPlayers int, winner string) {
	t.Helper()
	if got.StartedWith != numPlayers {
		t.Errorf("got %d players, wanted %d", got.StartedWith, numPlayers)
	}

	if got.FinishedWith != winner {
		t.Errorf("got winner %q, wanted %q", got.FinishedWith, winner)
	}
}

func AssertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout, wanted %q", got, messages)
	}
}

func AssertGameNotFinished(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.FinishCalled {
		t.Errorf("game shouldnt have finished")
	}
}