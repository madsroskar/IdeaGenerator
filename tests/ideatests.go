package main

import (
	"testing"
	"fmt"
)

func testTrimStrings(t *testing.T) {
	controlString := "Test"
	
	testString := fmt.Sprintf("\n%s\n", controlString)

	aRes, bRes := trimStrings(testString, testString)

	if aRes != controlString {
		t.Errorf("First result was incorrect, got: %s, want: %s.", aRes, controlString)
	}

	if aRes != controlString {
		t.Errorf("First result was incorrect, got: %s, want: %s.", aRes, controlString)
	}
}