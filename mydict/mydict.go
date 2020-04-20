package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string // 타입을 가명으로 바꿈

var errorNotFound = errors.New("Not Found")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}
