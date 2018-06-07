package main

import (
	"testing"
	"fmt"
	"strings"
	"math/rand"
)

func TestTrimStrings(t *testing.T) {
	controlString := "Test"
	testString := fmt.Sprintf("\n%s\n", controlString)

	aRes, bRes := trimStrings(testString, testString)

	if aRes != controlString {
		t.Errorf("First result was incorrect, got: %s, want: %s.", strings.Replace(aRes, "\n", "\\n", -1), controlString)
	}

	if bRes != controlString {
		t.Errorf("Second result was incorrect, got: %s, want: %s.", strings.Replace(bRes, "\n", "\\\\n", -1), controlString)
	}
}

func assertLen(o, r []string, t *testing.T) {
	if len(r) != len(o) {
        t.Errorf("Incorrect number of elements, got: %d, want: %d.", len(r), len(o))
    }
}

func TestSplitStrings(t *testing.T) {
	controlSlice := []string{ "a", "b", "c" }
	controlSliceFlat := strings.Trim(strings.Replace(fmt.Sprint(controlSlice), " ", ",", -1), "[]")
	testString := fmt.Sprintf("%s\n%s\n%s", "a", "b", "c")

	aRes, bRes := splitStrings(testString, testString)

	assertLen(controlSlice, aRes, t)

    aResFlat := strings.Trim(strings.Replace(fmt.Sprint(aRes), " ", ",", -1), "[]")

    if aResFlat != controlSliceFlat {
    	t.Errorf("First result was incorrect, got: %s, want: %s.", aResFlat, controlSliceFlat)
    }

	assertLen(controlSlice, bRes, t)

    bResFlat := strings.Trim(strings.Replace(fmt.Sprint(bRes), " ", ",", -1), "[]")

    if bResFlat != controlSliceFlat {
    	t.Errorf("First result was incorrect, got: %s, want: %s.", bResFlat, controlSliceFlat)
    }
    
}

func TestConvertByteArrays(t *testing.T) {
	testSlice := []byte("a\nb\nc")
	controlString := "a,b,c";

	aRes, bRes := convertByteArrays(testSlice, testSlice);
	aResFlat := strings.Trim(strings.Replace(fmt.Sprint(aRes), " ", ",", -1), "[]")
	bResFlat := strings.Trim(strings.Replace(fmt.Sprint(bRes), " ", ",", -1), "[]")

	if aResFlat != controlString {
    	t.Errorf("First result was incorrect, got: %s, want: %s.", aResFlat, controlString)
    }

    if bResFlat != controlString {
    	t.Errorf("First result was incorrect, got: %s, want: %s.", bResFlat, controlString)
    }
}

func TestIdeaString(t *testing.T) {
	idea, err := createIdea(sources, targets, 0, 0);
	controlString := "Spotify, but for Blockchain"

	if err != nil {
		t.Errorf("An error was raised: %s", err)
	}

	if idea.String() != controlString {
		t.Errorf("String was incorrect, got: %s, want: %s.", idea.String(), controlString)
	}
}

func TestCreateRandomIdea(t *testing.T) {
	rand.Seed(1)
	ideaOne := createRandomIdea(sources, targets);
	rand.Seed(1)
	ideaTwo := createRandomIdea(sources, targets);

	if ideaOne != ideaTwo {
		t.Errorf("Ideas not identical, got: %s, and: %s.", ideaOne.String(), ideaTwo.String())
	}
}

func TestGetValueByIndex(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}

	valZero, err := getValueByIndex(0, slice)

	if err != nil {
		t.Errorf("An error was raised: %s", err)
	}

	if valZero != "a" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", valZero, "a")
	}

	valTwo, err := getValueByIndex(2, slice)

	if err != nil {
		t.Errorf("An error was raised: %s", err)
	}

	if valTwo != "c" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", valTwo, "c")
	}

}

func TestCreateIdea(t *testing.T) {
	controlIdea := Idea{"Spotify", "Blockchain", 0, 0}
	res, err := createIdea(sources, targets, 0, 0)

	if err != nil {
		t.Errorf("An error was raised: %s", err)
	}

	if res != controlIdea {
		t.Errorf("Idea was not correct, got: %s, and: %s.", res.String(), controlIdea.String())
	}
}

