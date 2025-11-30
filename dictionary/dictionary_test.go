package dictionary_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/dictionary"
)

func assertStringsEqual(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("wanted %q got %q", want, got)
	}
}

func TestDictionary(t *testing.T) {
	test_dictionary := dictionary.Dictionary{"test": "this is just a test"}
	t.Run("existing word", func(t *testing.T) {
		got, _ := test_dictionary.Search("test")
		want := "this is just a test"

		assertStringsEqual(t, got, want)

	})

	t.Run("missing word", func(t *testing.T) {
		_, err := test_dictionary.Search("nonexisting")
		if err == nil {
			t.Error("expected an error but got none")
		}
	})

}
