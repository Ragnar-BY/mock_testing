package main

import (
	"testing"

	"github.com/gojuno/minimock"
)

var wrongkey = "WRONGKEY"

func TestDBProviderMinimock_ReadValue(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	db := NewDatabaseMock(mc)
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
		db.ReadMock.Expect(tc.key).Return(tc.expected, tc.err)

		val, err := dp.ReadValue(tc.key)
		if err != tc.err {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.err, err)
		}
		if val != tc.expected {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expected, val)
		}
	}
}

func TestDBProviderMinimock_AddValue(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	db := NewDatabaseMock(mc)
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
		db.WriteMock.Expect(tc.key, "val").Return(tc.err)
		err := dp.AddValue(tc.key, "val")
		if err != tc.err {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.err, err)
		}
	}
}
