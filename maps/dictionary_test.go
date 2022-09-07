package main

import "testing"

func TestSearch(t *testing.T) {
	
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("testing known words", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		AssertStrings(t, got, want, "test")
	})

	t.Run("testing unknown words", func(t *testing.T) {
		_, err := dictionary.Search("bob")
		
		AssertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "jo"
		definition := "a cool guy"
		
		err := dict.Add(word, definition)

		AssertError(t ,err, nil)
		AssertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "bob"
		definition := "sweetheart"
		dict := Dictionary{word: definition}
		err := dict.Add(word, "loser")

		AssertError(t, err, ErrorAlreadyExists)
		AssertDefinition(t, dict, word, definition)
	})	
}

func TestUpdate(t *testing.T) {
	
	t.Run("existing word", func(t *testing.T) {
		word := "bing"
		definition := "bong"
		newDefinition := "BONG"
		dict := Dictionary{word: definition}
		err := dict.Update(word, newDefinition)
	
		AssertError(t, err, nil)
		AssertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func (t *testing.T) {
		word := "ding"
		newDefinition := "DONG"
		dict := Dictionary{}
		err := dict.Update(word, newDefinition)

		AssertError(t, err, ErrorWordDoesntExist)
	})
}

func TestDelete(t *testing.T) {
	
	t.Run("existing word", func(t *testing.T) {
		word := "bing"
		dict := Dictionary{word: "bong"}
		dict.Delete(word)
	
		_, err := dict.Search(word)
		if err != ErrorNotFound {
			t.Errorf("%q should have been deleted but wasnt", word)
		}
	})
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func AssertStrings(t testing.TB, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q, given %q", got, want, given)
	}
}

func AssertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}