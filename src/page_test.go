package main

import (
	"testing"
	"reflect"
)

func TestFormatJsonDataIdea(t *testing.T) {
	idea := createRandomIdea(sources, targets)
	page := Page{idea, ""}

	res := page.formatJsonData()
	control := map[string]string{"Idea": idea.String()}

	switch res.(type) {
	case map[string]string:
	default:
		t.Errorf("Type mismatch. Got %v, expected map[string]string", res)
	}

	if !reflect.DeepEqual(res, control) {
		t.Error("The result doesn't match expected value.")
	}
}

func TestFormatJsonDataMap(t *testing.T) {
	idea := createRandomIdea(sources, targets)
	tmap := map[string]Idea{"Idea": idea}
	page := Page{tmap, ""}

	res := page.formatJsonData()

	switch res.(type) {
	case map[string]Idea:
	default:
		t.Errorf("Type mismatch. Got %T, expected map[string]Idea", res)
	}

	if !reflect.DeepEqual(res, tmap) {
		t.Error("The result doesn't match expected value.")
	}	
}
