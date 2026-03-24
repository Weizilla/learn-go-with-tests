package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}

		got, _ := dict.Search("test")
		want := "this is a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	t.Run("unkonwn word", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}

		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expect to get error")
		}
		if err != ErrorNotFound {
			t.Fatal("wrong error")
		}

	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("test", "this is a test")

		want := "this is a test"
		got, _ := dict.Search("test")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		value := "this is just a test"
		dict := Dictionary{word: value}

		err := dict.Add(word, "new test")

		if err == nil {
			t.Fatal("expect to get error")
		}
		if err != ErrWordExists {
			t.Fatal("wrong error")
		}
		got, _ := dict.Search("test")

		if got != value {
			t.Errorf("got %q want %q given, %q", got, value, "test")
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		value := "this is just a test"
		dict := Dictionary{word: value}

		newValue := "this is a new value"
		dict.Update(word, newValue)

		got, _ := dict.Search("test")

		if got != newValue {
			t.Errorf("got %q want %q given, %q", got, newValue, "test")
		}
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		value := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, value)

		if err == nil {
			t.Fatal("expect to get error")
		}
		if err != ErrWordDoesNotExist {
			t.Fatal("wrong error")
		}

	})
}

func TestName(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "value"}

		dict.Delete(word)

		_, err := dict.Search(word)

		if err == nil {
			t.Fatal("expect to get error")
		}
		if err != ErrorNotFound {
			t.Fatal("wrong error")
		}
	})
}
