package main

import (
	"testing"
)

var wrongkey = "WRONGKEY"

func TestDBProviderMoq_ReadValue(t *testing.T) {
	db := &DatabaseMock{ReadFunc: func(key string) (string, error) {
		if key == wrongkey {
			return "", ErrWrongKey
		}
		return "value", nil
	}}
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
		val, err := dp.ReadValue(tc.key)
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedError, err)
		}
		if val != tc.expectedValue {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedValue, val)
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
		err := dp.AddValue(tc.key, "val")
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedError, err)
		}
	}
}
