package main

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrorNotFound	     = DictionaryErr("couldn't find the word")
	ErrorAlreadyExists   = DictionaryErr("this word already exists")
	ErrorWordDoesntExist = DictionaryErr("cant update, word doesn't exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]
	
	if !exists {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorAlreadyExists
	default:
		return err
	}

	return nil 
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesntExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}