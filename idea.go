package main

import (
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

func createIdea(sources, targets []byte, source, target int) Idea {
	s, t := convertByteArrays(sources, targets)

	return Idea{s[source], t[target], source, target}
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
