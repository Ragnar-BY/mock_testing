package main

import (
	"testing"
)

var wrongkey = "WRONGKEY"

func TestDBProviderCounterFeiter_ReadValue(t *testing.T) {
	db := new(FakeDatabase)
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
		db.ReadReturns(tc.expected, tc.err)
		val, err := dp.ReadValue(tc.key)
		if err != tc.err {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.err, err)
		}
		if val != tc.expected {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expected, val)
		}
	}
}

func TestDBProviderCounterFeiter_AddValue(t *testing.T) {
	db := new(FakeDatabase)
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
		db.WriteReturns(tc.err)
		err := dp.AddValue(tc.key, "val")
		if err != tc.err {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.err, err)
		}
	}
}
