package dictionary

type errorConst string

const (
	ErrWordMissing errorConst = "word not found"
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
