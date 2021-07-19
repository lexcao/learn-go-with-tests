package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTempFile(t testing.TB, data string) (*os.File, func()) {
	t.Helper()

	tempFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, _ = tempFile.WriteString(data)

	removeFile := func() {
		_ = tempFile.Close()
		_ = os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database, clean := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer clean()

		store := NewFileSystemPlayerStore(database)
		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, clean := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer clean()

		store := NewFileSystemPlayerStore(database)
		got := store.GetPlayerScore("Chris")

		want := 33

		assertScore(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, clean := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer clean()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScore(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, clean := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer clean()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScore(t, got, want)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, clean := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
		defer clean()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

func assertScore(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
