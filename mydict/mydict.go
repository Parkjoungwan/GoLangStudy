package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string // 타입을 가명으로 바꿈

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word alreay exists")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
