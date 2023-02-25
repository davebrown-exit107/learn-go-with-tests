package dictionary

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("the word is already defined in the dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, present := d[word]
	if present {
		return ErrWordExists
	} else {
		d[word] = definition
		return nil
	}

}
