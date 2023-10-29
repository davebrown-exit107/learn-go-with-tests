package dictionary_test

import (
	"testing"

	dictionary "github.com/davebrown-exit107/learn-go-with-tests/maps"
)

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
		t.Errorf("wanted %v got %v", want, got)
	}
}

func assertStringsEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q got %q", want, got)
	}
}
