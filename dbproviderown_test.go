package main

import (
	"testing"
)

// We create own mock implementation of DB.
type dbmock struct{}

var wrongkey = "WRONGKEY"

func (db *dbmock) Read(key string) (string, error) {
	if key == wrongkey {
		return "", ErrWrongKey
	}
	return "value", nil
}
func (db *dbmock) Write(key string, value string) error {
	if key == wrongkey {
		return ErrWrongKey
	}
	return nil
}
func TestDBProvider_ReadValue(t *testing.T) {
	db := &dbmock{}
	dp := DBProvider{db}

	tt := []struct {
		name     string
		key      string
		expected string
		err      error
	}{
		{
			name:     "right key",
			key:      "rightkey",
			expected: "value",
			err:      nil,
		},
		{
			name:     "wrong key",
			key:      wrongkey,
			expected: "",
			err:      ErrWrongKey,
		},
	}

	for _, tc := range tt {
		val, err := dp.ReadValue(tc.key)
		if err != tc.err {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.err, err)
		}
		if val != tc.expected {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expected, val)
		}
	}
}

func TestDBProvider_AddValue(t *testing.T) {
	db := &dbmock{}
	dp := DBProvider{db}

	tt := []struct {
		name string
		key  string
		err  error
	}{
		{
			name: "right key",
			key:  "rightkey",
			err:  nil,
		},
		{
			name: "wrong key",
			key:  wrongkey,
			err:  ErrWrongKey,
		},
	}

	for _, tc := range tt {
		err := dp.AddValue(tc.key, "val")
		if err != tc.err {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.err, err)
		}
	}
}
