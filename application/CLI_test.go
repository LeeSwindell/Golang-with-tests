package poker_test

import (
	poker "application"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	
	t.Run("testing chris", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
	
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
	
		assertPlayerWins(t, playerStore, "Chris")
	})

	t.Run("testing cleo", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
	
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
	
		assertPlayerWins(t, playerStore, "Cleo")
	})
}

func assertPlayerWins(t testing.TB, store *poker.StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatal("expected a win and didnt get one")
	}

	if store.WinCalls[0] != winner {
		t.Errorf("got winner %q, wanted %q", store.WinCalls[0], winner)
	}
}