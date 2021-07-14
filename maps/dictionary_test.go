package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := "could not find the word you were looking for"

		assertError(t, err, ErrNotFound)

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("new word", func(t *testing.T) {
		dict := Dict{}
		err := dict.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dict{word: definition}
		err := dict.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("existing word", func(t *testing.T) {
		dict := Dict{word: definition}
		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dict := Dict{}
		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dict := Dict{word: definition}

	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dict Dict, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	assertNoError(t, err)
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
