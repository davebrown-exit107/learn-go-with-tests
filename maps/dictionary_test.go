package dictionary_test

import (
	"testing"

	dictionary "github.com/davebrown-exit107/learn-go-with-tests/maps"
)

func TestSearch(t *testing.T) {
	testDictionary := dictionary.Dictionary{"test": "this is just a test"}

	got := testDictionary.Search("test")
	want := "this is just a test"

	assertStringsEqual(t, got, want)

}

func assertStringsEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q got %q", want, got)
	}
}
