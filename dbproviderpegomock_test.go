package main

import (
	"testing"

	"github.com/petergtz/pegomock"
)

var wrongkey = "WRONGKEY"

func TestDBProviderPegomock_ReadValue(t *testing.T) {

	pegomock.RegisterMockTestingT(t)

	db := NewMockDatabase()
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
		pegomock.When(db.Read(tc.key)).ThenReturn(tc.expected, tc.err)
		val, err := dp.ReadValue(tc.key)
		if err != tc.err {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.err, err)
		}
		if val != tc.expected {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expected, val)
		}
	}
}
func TestDBProviderPegomock_AddValue(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	db := NewMockDatabase()
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
		pegomock.When(db.Write(tc.key, "val")).ThenReturn(tc.err)
		err := dp.AddValue(tc.key, "val")
		if err != tc.err {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.err, err)
		}
	}
}
