package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var wrongkey = "WRONGKEY"

type dbmock struct {
	mock.Mock
}

func (db *dbmock) Read(key string) (string, error) {
	args := db.Called(key)
	return args.String(0), args.Error(1)
}
func (db *dbmock) Write(key string, value string) error {
	args := db.Called(key, value)
	return args.Error(0)

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
		db.On("Read", tc.key).Return(tc.expected, tc.err)
		val, err := dp.ReadValue(tc.key)
		assert.Equal(t, tc.expected, val)
		assert.Equal(t, tc.err, err)
		db.AssertExpectations(t)
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
		db.On("Write", tc.key, "val").Return(tc.err)
		err := dp.AddValue(tc.key, "val")
		assert.Equal(t, tc.err, err)
		db.AssertExpectations(t)
	}
}
