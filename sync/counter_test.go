package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := New()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter.Value(), 3)
	})

	t.Run("it runs safly concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := New()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter.Value(), wantedCount)
	})
}

func assertCounter(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
