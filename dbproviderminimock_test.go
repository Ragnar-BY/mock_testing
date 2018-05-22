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
		db.ReadMock.Expect(tc.key).Return(tc.expectedValue, tc.expectedError)
		val, err := dp.ReadValue(tc.key)
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedError, err)
		}
		if val != tc.expectedValue {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedValue, val)
		}
	}
}

func TestDBProviderMinimock_AddValue(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	db := NewDatabaseMock(mc)
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
		db.WriteMock.Expect(tc.key, "val").Return(tc.expectedError)
		err := dp.AddValue(tc.key, "val")
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedError, err)
		}
	}
}
