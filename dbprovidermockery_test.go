package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wrongkey = "WRONGKEY"

func TestDBProviderMockery_ReadValue(t *testing.T) {
	db := &MockDatabase{}
	dp := DBProvider{db}
	tt := []struct {
		name          string
		key           string
		expectedValue string
		expectedError error
	}{
		{
			name:          "right key",
			key:           "rightkey",
			expectedValue: "value",
			expectedError: nil,
		},
		{
			name:          "wrong key",
			key:           wrongkey,
			expectedValue: "",
			expectedError: ErrWrongKey,
		},
	}
	for _, tc := range tt {
		db.On("Read", tc.key).Return(tc.expectedValue, tc.expectedError)
	}
	for _, tc := range tt {
		val, err := dp.ReadValue(tc.key)
		assert.Equal(t, tc.expectedValue, val)
		assert.Equal(t, tc.expectedError, err)

	}
	db.AssertExpectations(t)
}

func TestDBProviderMockery_AddValue(t *testing.T) {
	db := &MockDatabase{}
	dp := DBProvider{db}
	tt := []struct {
		name          string
		key           string
		expectedError error
	}{
		{
			name:          "right key",
			key:           "rightkey",
			expectedError: nil,
		},
		{
			name:          "wrong key",
			key:           wrongkey,
			expectedError: ErrWrongKey,
		},
	}
	for _, tc := range tt {
		db.On("Write", tc.key, "val").Return(tc.expectedError)
	}
	for _, tc := range tt {
		err := dp.AddValue(tc.key, "val")
		assert.Equal(t, tc.expectedError, err)
	}
	db.AssertExpectations(t)
}
