package dictionary_test

import (
	"testing"

	dictionary "github.com/davebrown-exit107/learn-go-with-tests/maps"
)

func TestSearch(t *testing.T) {
	testDictionary := map[string]string{"test": "this is just a test"}

	got := dictionary.Search(testDictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("wanted %q got %q", want, got)
	}

}
