package main

import "errors"

//go:generate minimock -i github.com/Ragnar-BY/mock_testing.Database -o ./

// ErrWrongKey is error for wrong key.
var ErrWrongKey = errors.New("wrong key")

// Database is database interface.
type Database interface {
	Read(key string) (string, error)
	Write(key string, value string) error
}
