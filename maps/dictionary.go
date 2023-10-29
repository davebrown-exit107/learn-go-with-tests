package dictionary

type errorConst string

const (
	ErrWordMissing errorConst = "word not found"
	ErrWordExists  errorConst = "word already exists"
)

func (c errorConst) Error() string {
	return string(c)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if ok {
		return val, nil
	} else {
		return "", ErrWordMissing
	}
}

func (d Dictionary) Add(word, definition string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	} else {
		d[word] = definition
		return nil
	}
}
