package mock

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type CountdownOperationSpy struct {
	calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.calls = append(s.calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownOperationSpy{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spy := &CountdownOperationSpy{}
		Countdown(spy, spy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spy.calls) {
			t.Errorf("wanted calls %v got %v", want, spy.calls)
		}
	})
}

type SpyTime struct {
	duration time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.duration = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.duration != sleepTime {
		t.Errorf("should have slept for %v but slept for %b", sleepTime, spyTime)
	}
}
