package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1,2,3}, []int{4,5})
	want := []int{6,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("testing normal sized slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{5, 6})
		want := []int{5, 6}
		checkSums(t, got, want)
	})
	
	t.Run("testing empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1,2,3})
		want := []int{0,5}
		checkSums(t, got, want)
	})

}