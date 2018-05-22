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

func TestDBProvider_AddValue(t *testing.T) {
	db := &dbmock{}
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
