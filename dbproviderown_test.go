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
		testName      string
		key           string
		expectedValue string
		expectedError error
	}{
		{
			testName:      "right key",
			key:           "rightkey",
			expectedValue: "value",
			expectedError: nil,
		},
		{
			testName:      "wrong key",
			key:           wrongkey,
			expectedValue: "",
			expectedError: ErrWrongKey,
		},
	}

	for _, tc := range tt {
		val, err := dp.ReadValue(tc.key)
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.testName, tc.expectedError, err)
		}
		if val != tc.expectedValue {
			t.Errorf("[%s] expected %v, received %v", tc.testName, tc.expectedValue, val)
		}
	}
}

func TestDBProvider_AddValue(t *testing.T) {
	db := &dbmock{}
	dp := DBProvider{db}

	tt := []struct {
		testName      string
		key           string
		expectedError error
	}{
		{
			testName:      "right key",
			key:           "rightkey",
			expectedError: nil,
		},
		{
			testName:      "wrong key",
			key:           wrongkey,
			expectedError: ErrWrongKey,
		},
	}

	for _, tc := range tt {
		err := dp.AddValue(tc.key, "val")
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.testName, tc.expectedError, err)
		}
	}
}
