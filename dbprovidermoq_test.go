package main

import (
	"testing"
)

// Tests use mock from github.com/matryer/moq
func TestDBProviderMoq_ReadValue(t *testing.T) {
	db := &DatabaseMock{ReadFunc: func(key string) (string, error) {
		if key == wrongkey {
			return "", ErrWrongKey
		}
		return "value", nil
	}}
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

func TestDBProviderMoq_AddValue(t *testing.T) {
	db := &DatabaseMock{WriteFunc: func(key string, value string) error {
		if key == wrongkey {
			return ErrWrongKey
		}
		return nil
	}}
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
