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
		pegomock.When(db.Read(tc.key)).ThenReturn(tc.expectedValue, tc.expectedError)
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
func TestDBProviderPegomock_AddValue(t *testing.T) {
	pegomock.RegisterMockTestingT(t)
	db := NewMockDatabase()
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
		pegomock.When(db.Write(tc.key, "val")).ThenReturn(tc.expectedError)
	}
	for _, tc := range tt {
		err := dp.AddValue(tc.key, "val")
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedError, err)
		}
	}
}
