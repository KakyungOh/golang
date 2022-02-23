package mydict

import "errors"

// Dictionary
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCanUdate   = errors.New("Can't update non-existing word")
	errCanDelete  = errors.New("Can't delete non-existing word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCanUdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCanDelete
	}
	return nil
}
