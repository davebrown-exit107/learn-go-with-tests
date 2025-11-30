package dictionary

import "errors"

type Dictionary map[string]string

var ErrWordNotFound = errors.New("word missing from dictionary")

func (d Dictionary) Search(search_term string) (string, error) {
	definition, ok := d[search_term]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}
