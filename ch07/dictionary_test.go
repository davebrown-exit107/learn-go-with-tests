package dictionary_test

import (
	"testing"

	dictionary "github.com/davebrown-exit107/learn-go-with-tests/ch07"
)

func TestDelete(t *testing.T) {
	t.Run("delete word in dictionary", func(t *testing.T) {
		testDictionary := dictionary.Dictionary{"test": "this is just a test"}
		testDictionary.Delete("test")

		_, got := testDictionary.Search("test")
		want := dictionary.ErrWordMissing
		assertError(t, got, want)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("update word in dictionary", func(t *testing.T) {
		testDictionary := dictionary.Dictionary{"test": "this is just a test"}

		want := "a new definition"
		_ = testDictionary.Update("test", want)

		got, _ := testDictionary.Search("test")
		assertStringsEqual(t, want, got)
	})
	t.Run("update missing word in dictionary", func(t *testing.T) {
		testDictionary := dictionary.Dictionary{}

		got := testDictionary.Update("another", "a new definition")
		want := dictionary.ErrWordDoesNotExist

		assertError(t, want, got)
	})
}
func TestAdd(t *testing.T) {
	t.Run("add word to dictionary", func(t *testing.T) {
		testDictionary := dictionary.Dictionary{}

		want := "a new word"
		_ = testDictionary.Add("new", want)

		got, _ := testDictionary.Search("new")
		assertStringsEqual(t, got, want)
	})
	t.Run("add same word twice", func(t *testing.T) {
		testDictionary := dictionary.Dictionary{"new": "a new word"}
		got := testDictionary.Add("new", "a new word")
		want := dictionary.ErrWordExists
		assertError(t, got, want)
	})

}

func TestSearch(t *testing.T) {

	t.Run("search for word not in dictionary", func(t *testing.T) {

		testDictionary := dictionary.Dictionary{"test": "this is just a test"}

		_, err := testDictionary.Search("unknown")
		want := dictionary.ErrWordMissing

		assertError(t, err, want)

	})

	t.Run("search for word in dictionary", func(t *testing.T) {

		testDictionary := dictionary.Dictionary{"test": "this is just a test"}

		got, _ := testDictionary.Search("test")
		want := "this is just a test"

		assertStringsEqual(t, got, want)

	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q got %q", want, got)
	}
}

func assertStringsEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q got %q", want, got)
	}
}
