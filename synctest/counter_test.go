package synctest

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("works concurrently", func(t *testing.T) {
		wantCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, wantCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}