package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "value"}

	t.Run("search", func(t *testing.T) {

		actual, _ := dict.Search("test")
		expected := "value"

		assert.Equal(t, expected, actual)
	})

	t.Run("search no result", func(t *testing.T) {
		_, err := dict.Search("something")

		assert.ErrorIs(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Add("key", "value")
		assert.Nil(t, err)

		got, err := dict.Search("key")

		assert.Nil(t, err)
		assert.Equal(t, got, "value")
	})

	t.Run("add existing", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Add("key", "value")
		assert.Nil(t, err)

		err = dict.Add("key", "value")
		assert.ErrorIs(t, err, ErrExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update", func(t *testing.T) {
		value := "value"
		key := "key"
		newValue := "new value"
		dict := Dictionary{key: value}

		_ = dict.Update(key, newValue)

		search, _ := dict.Search(key)
		assert.Equal(t, newValue, search)
	})

	t.Run("no update", func(t *testing.T) {
		key := "key"
		value := "value"

		dict := Dictionary{}

		err := dict.Update(key, value)

		assert.ErrorIs(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		key := "key"
		value := "value"
		dict := Dictionary{key: value}

		dict.Delete(key)

		_, err := dict.Search(key)
		assert.ErrorIs(t, err, ErrNotFound)
	})
}
