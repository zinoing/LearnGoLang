package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("not Found")
	errWordExists = errors.New("there is already word in dictionary")
	errCantUpdate = errors.New("cant update non-existing word")
	errCantDelete = errors.New("cant delete non-existing word")
)

// Search for words
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		return errWordExists
	}
	d[word] = def
	return nil
}

// Update word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = def
	} else {
		return errCantUpdate
	}
	return nil
}

// Delete word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == nil {
		delete(d, word)
	} else {
		return errCantDelete
	}
	return nil
}
