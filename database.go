package main

import "errors"

//go:generate counterfeiter -o ./database_mock.go . Database

// ErrWrongKey is error for wrong key.
var ErrWrongKey = errors.New("wrong key")

// Database is database interface.
type Database interface {
	Read(key string) (string, error)
	Write(key string, value string) error
}
