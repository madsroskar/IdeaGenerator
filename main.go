package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

type Idea struct {
	source string
	target string
}

func (i Idea) String() string {
	return fmt.Sprintf("%s, but for %s", i.source, i.target)
}

func createIdea(sources, targets []byte) Idea {
	trimmedSources := strings.TrimSpace(string(sources))
	trimmedTargets := strings.TrimSpace(string(targets))
	
	s := strings.Split(trimmedSources, "\n")
	t := strings.Split(trimmedTargets, "\n")
	
	return Idea{s[random(0, len(s))], t[random(0, len(t))]}
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

func main() {
	rand.Seed(time.Now().Unix())
	var sources, targets []byte = readFiles("existing-idea.txt", "targets.txt")
	idea := createIdea(sources, targets)
	fmt.Println(idea)
}

