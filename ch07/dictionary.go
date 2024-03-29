package dictionary

type errorConst string

const (
	ErrWordMissing      errorConst = "word not found"
	ErrWordDoesNotExist errorConst = "could not update, word not found"
	ErrWordExists       errorConst = "word already exists"
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

func (d Dictionary) Update(word, definition string) error {
	if _, ok := d[word]; ok {
		d[word] = definition
		return nil
	} else {
		return ErrWordDoesNotExist
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
