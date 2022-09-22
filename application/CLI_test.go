package poker_test

import (
	poker "application"
	"bytes"
	"io"
	"strings"
	"testing"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	
	t.Run("testing chris", func(t *testing.T) {
		in := userSends("3", "Chris wins")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}
	
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		poker.AssertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertCorrectGame(t, game, 3, "Chris")	
	})

	t.Run("testing cleo", func(t *testing.T) {
		in := userSends("1", "Cleo wins")
		game := &poker.GameSpy{}
	
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
	
		poker.AssertCorrectGame(t, game, 1, "Cleo")
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &poker.GameSpy{FinishCalled: false}
		stdout := &bytes.Buffer{}

		in := userSends("8", "Lloyd is a killer")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		poker.AssertGameNotFinished(t, game)
		poker.AssertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)
	})

	t.Run("error when a non numeric value is entered, does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("Pies")
		game := &poker.GameSpy{StartCalled: false}
	
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		
		poker.AssertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}
