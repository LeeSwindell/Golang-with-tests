package generics2

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {

	t.Run("testing Sum", func(t *testing.T) {		
		got := Sum([]int{1,4,5})
		want := 10

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("testing SumTails", func(t *testing.T) {
		a, b := []int{1,2,3,4}, []int{1,2}
		got := SumAllTails(a, b)
		want := []int{9, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}


func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}