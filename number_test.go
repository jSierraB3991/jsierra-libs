package eliotlibs

import (
	"testing"
)

func TestConvertFloat32ToString(t *testing.T) {

	testCases := []struct {
		Value       float32
		ExpectValue string
	}{
		{
			Value:       123.45,
			ExpectValue: "123.45",
		},
		{
			Value:       595.34,
			ExpectValue: "595.34",
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.ExpectValue, func(t *testing.T) {
			result := GetStrinToFloat(tc.Value)
			if result != tc.ExpectValue {
				t.Errorf("expect %v, got %v", tc.ExpectValue, result)
			}
		})
	}
}

func TestConvertStringToFloat32(t *testing.T) {

	testCases := []struct {
		Value       string
		ExpectValue float32
		ExpectError error
	}{
		{
			Value:       "123.45",
			ExpectValue: 123.45,
			ExpectError: nil,
		},
		{
			Value:       "595.34",
			ExpectValue: 595.34,
			ExpectError: nil,
		},
		{
			Value:       "fg",
			ExpectValue: 0,
			ExpectError: InvalidConvertStringToFloat{},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Value, func(t *testing.T) {
			result, err := GetStringtofloat32(tc.Value)
			if err != tc.ExpectError {
				t.Errorf("expect %v, got %v", tc.ExpectError, err)
			}
			if result != tc.ExpectValue {
				t.Errorf("expect %v, got %v", tc.ExpectValue, result)
			}
		})
	}
}

func TestConvertFloat32ToFomatMoney(t *testing.T) {

	testCases := []struct {
		Value       float64
		ExpectValue string
		ExpectError error
	}{
		{
			Value:       1123.45,
			ExpectValue: "$1,123.45",
			ExpectError: nil,
		},
		{
			Value:       595.34,
			ExpectValue: "$595.34",
			ExpectError: nil,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.ExpectValue, func(t *testing.T) {
			result := GetFormatMoneyToFloat(tc.Value)
			if result != tc.ExpectValue {
				t.Errorf("expect %v, got %v", tc.ExpectValue, result)
			}
		})
	}
}
