package dict_test

import (
	"testing"

	dict "github.com/davebrown-exit107/learn-go-with-tests/dictionary"
)

func TestSearch(t *testing.T) {
	dictionary := dict.Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "unknown"
		_, err := dictionary.Search(word)

		assertError(t, err, dict.ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := dict.Dictionary{}
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("dupe word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := dict.Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, dict.ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := dict.Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := dict.Dictionary{word: definition}
		newWord := "new definition"

		err := dictionary.Update(newWord, definition)

		assertError(t, err, dict.ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := dict.Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, dict.ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary dict.Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
