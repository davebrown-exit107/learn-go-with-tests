package dictionary_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/dictionary"
)

func TestDictionary(t *testing.T) {
	test_dictionary := map[string]string{"test": "this is just a test"}

	got := dictionary.Search("test", test_dictionary)
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q wanted %q, given %q", got, want, "test")
	}
}
