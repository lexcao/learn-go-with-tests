package reader

import (
	"context"
	"io"
	"strings"
	"testing"
)

func TestContextAwareReader(t *testing.T) {
	t.Run("let's just see how a normal reader works", func(t *testing.T) {
		reader := strings.NewReader("123456")
		assertRead(t, reader, "123")
		assertRead(t, reader, "456")
	})

	t.Run("behaves like a normal reader", func(t *testing.T) {
		reader := NewCancellableReader(context.Background(), strings.NewReader("123456"))
		assertRead(t, reader, "123")
		assertRead(t, reader, "456")
	})

	t.Run("stops reading when cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		reader := NewCancellableReader(ctx, strings.NewReader("123456"))
		assertRead(t, reader, "123")

		cancel()

		got := make([]byte, 3)
		n, err := reader.Read(got)

		if err == nil {
			t.Error("expected an error, but got none")
		}

		if n > 0 {
			t.Errorf("expected 0 bytes to be read, but got %d", n)
		}
	})
}

func assertRead(t testing.TB, reader io.Reader, want string) {
	t.Helper()
	got := make([]byte, len(want))
	n, err := reader.Read(got)

	if n != len(want) {
		t.Errorf("got %d bytes read, want %d", n, len(want))
	}

	if err != nil {
		t.Fatal(err)
	}

	assertBufferHas(t, got, want)
}

func assertBufferHas(t testing.TB, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
