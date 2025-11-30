package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(search_term string) (string, error) {
	definition, ok := d[search_term]
	if !ok {
		return "", errors.New("word missing from dictionary")
	}
	return definition, nil
}
