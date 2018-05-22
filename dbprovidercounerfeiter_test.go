package main

import (
	"testing"
)

var wrongkey = "WRONGKEY"

func TestDBProviderCounterFeiter_ReadValue(t *testing.T) {
	db := new(FakeDatabase)
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
		db.ReadReturns(tc.expectedValue, tc.expectedError)
		val, err := dp.ReadValue(tc.key)
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedError, err)
		}
		if val != tc.expectedValue {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedValue, val)
		}
	}
}

func TestDBProviderCounterFeiter_AddValue(t *testing.T) {
	db := new(FakeDatabase)
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
		db.WriteReturns(tc.expectedError)
		err := dp.AddValue(tc.key, "val")
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedError, err)
		}
	}
}
