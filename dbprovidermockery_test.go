package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wrongkey = "WRONGKEY"

func TestDBProviderMockery_ReadValue(t *testing.T) {
	db := &MockDatabase{}
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
		dp := DBProvider{db}
		val, err := dp.ReadValue(tc.key)
		assert.Equal(t, tc.expected, val)
		assert.Equal(t, tc.err, err)
		db.AssertExpectations(t)
	}
}

func TestDBProviderMockery_AddValue(t *testing.T) {
	db := &MockDatabase{}
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
		dp := DBProvider{db}
		err := dp.AddValue(tc.key, "val")
		assert.Equal(t, tc.err, err)
		db.AssertExpectations(t)
	}
}
