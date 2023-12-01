package maps

type Dictionary map[string]string

const (
	ErrNotFound = DictErr("not found")
	ErrExists   = DictErr("exists")
)

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
