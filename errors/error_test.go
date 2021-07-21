package errors

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func DumbGetter(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("fetching from %s, %v", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return "", BadStatusError{url, res.StatusCode}
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	return string(body), nil
}

type BadStatusError struct {
	URL    string
	Status int
}

func (b BadStatusError) Error() string {
	return fmt.Sprintf("not 200 from %s, got %d", b.URL, b.Status)
}

func TestDumbGetter(t *testing.T) {
	t.Run("not 200", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		}))
		defer server.Close()

		_, err := DumbGetter(server.URL)

		if err == nil {
			t.Fatalf("expected an error")
		}

		var got BadStatusError
		isBadStatusErr := errors.As(err, &got)
		if !isBadStatusErr {
			t.Fatalf("want BadStatusErr, got %T", err)
		}

		want := BadStatusError{server.URL, http.StatusTeapot}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
