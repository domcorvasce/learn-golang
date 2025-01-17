package main

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

func (dict Dictionary) Search(term string) (string, error) {
	value, ok := dict[term]
	if !ok {
		return value, errors.New(fmt.Sprintf("cannot retrieve value for key '%s'", term))
	}

	return value, nil
}

func (dict Dictionary) Add(term string, definition string) error {
	_, err := dict.Search(term)
	if err == nil {
		return errors.New("cannot overwrite existing term")
	}

	dict[term] = definition
	return nil
}
