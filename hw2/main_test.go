package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	testData := [][]string{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"}}
	for slice := range testData {
		result, unpackError := Unpack(testData[slice][0])
		if unpackError != nil {
			t.Fatalf("test failed with error for input %s", testData[slice][0])
		}
		if result != testData[slice][1] {
			t.Fatalf("invalid string returned for input %s", testData[slice][0])
		}
	}


}
