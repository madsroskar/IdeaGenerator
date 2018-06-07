package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Idea struct {
	source   string
	target   string
	SourceId int
	TargetId int
}

func trimStrings(a, b string) (string, string) {
	return strings.TrimSpace(a), strings.TrimSpace(b)
}

func splitStrings(a, b string) ([]string, []string) {
	return strings.Split(a, "\n"), strings.Split(b, "\n")
}

func convertByteArrays(a, b []byte) ([]string, []string) {
	atrim, btrim := trimStrings(string(a), string(b))
	return splitStrings(atrim, btrim)
}

func (i Idea) String() string {
	return fmt.Sprintf("%s, but for %s", i.source, i.target)
}

func createRandomIdea(sources, targets []byte) Idea {
	s, t := convertByteArrays(sources, targets)

	source := random(0, len(s))
	target := random(0, len(t))

	return Idea{s[source], t[target], source, target}
}

func getValueByIndex(i int, s []string) (string, error) {
	if i < 0 || i > len(s)-1 {
		return "", errors.New("Index is out of bounds")
	}
	return s[i], nil
}

func createIdea(sources, targets []byte, sourceIndex, targetIndex int) (Idea, error) {
	s, t := convertByteArrays(sources, targets)

	source, serr := getValueByIndex(sourceIndex, s)
	target, terr := getValueByIndex(targetIndex, t)

	if serr != nil || terr != nil {
		return Idea{}, errors.New("Idea not found.")
	}

	return Idea{source, target, sourceIndex, targetIndex}, nil
}

func readFile(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	return content
}

func readFiles(sourceFileName, targetFileName string) ([]byte, []byte) {
	return readFile(sourceFileName), readFile(targetFileName)
}
