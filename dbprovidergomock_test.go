package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

var wrongkey = "WRONGKEY"

func TestDBProviderGomock_ReadValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := NewMockDatabase(ctrl)
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
		db.EXPECT().Read(tc.key).Return(tc.expectedValue, tc.expectedError)
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

func TestDBProviderGomock_AddValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := NewMockDatabase(ctrl)
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
		db.EXPECT().Write(tc.key, "val").Return(tc.expectedError)
	}
	for _, tc := range tt {
		err := dp.AddValue(tc.key, "val")
		if err != tc.expectedError {
			t.Errorf("[%s] expected %v, received %v", tc.name, tc.expectedError, err)
		}
	}
}
