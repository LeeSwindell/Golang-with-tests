package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on ints", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "abc", "abc")
		AssertNotEqual(t, "123", "abc")
	})

	// AssertEqual(t, 1, "1")
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("%+v is not equal to %+v!!", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("%+v is equal to %+v!!", got, want)
	}
}

func TestStackofInts(t *testing.T) {
	
	t.Run("testing push", func(t *testing.T) {
		got := Stack[int]{values: []int{}}
		got.Push(1)
		want := Stack[int]{values: []int{1}}

		AssertEqual(t, got.values[0], want.values[0])
	})

	t.Run("testing pop", func(t *testing.T) {
		stack := Stack[int]{values: []int{1}}
		got, _ := stack.Pop()
		want := 1

		AssertEqual(t, got, want)
		AssertEqual(t, len(stack.values), 0)
	})

	t.Run("testing isEmpty", func(t *testing.T) {
		stack := Stack[int]{values: []int{}}
		got := stack.isEmpty()
		want := true
		AssertEqual(t, got, want)

		stack = Stack[int]{values: []int{1}}
		got = stack.isEmpty()
		want = false
		AssertEqual(t, got, want)
	})
}

func TestStackofStrings(t *testing.T) {
	
	t.Run("testing push", func(t *testing.T) {
		got := Stack[string]{values: []string{}}
		got.Push("1")
		want := Stack[string]{values: []string{"1"}}

		AssertEqual(t, got.values[0], want.values[0])
	})

	t.Run("testing pop", func(t *testing.T) {
		stack := Stack[string]{values: []string{"1"}}
		got, _ := stack.Pop()
		want := "1"

		AssertEqual(t, got, want)
		AssertEqual(t, len(stack.values), 0)
	})

	t.Run("testing isEmpty", func(t *testing.T) {
		stack := Stack[string]{values: []string{}}
		got := stack.isEmpty()
		want := true
		AssertEqual(t, got, want)

		stack = Stack[string]{values: []string{"1"}}
		got = stack.isEmpty()
		want = false
		AssertEqual(t, got, want)
	})
}