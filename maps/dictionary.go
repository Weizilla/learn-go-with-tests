package maps

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrorNotFound       = DictionaryError("not found")
	ErrWordExists       = DictionaryError("word exists")
	ErrWordDoesNotExist = DictionaryError("word does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		delete(d, key)
		return nil
	default:
		return err
	}
}
