package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("testing deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("testing withdrawal with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("testing overdrawn withdrawl", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
	
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}


func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance

	if wallet.balance != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("There was an error!")
	}
}