package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecond = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecond)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		winner = a
	case <-ping(b):
		winner = b
	case <-time.After(timeout):
		err = fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	return
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
