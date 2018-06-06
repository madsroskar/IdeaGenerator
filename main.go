package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
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
	s := strings.Split(string(sources), "\n")
	t := strings.Split(string(targets), "\n")
	return Idea{s[random(0, len(s) - 1)], t[random(0, len(t) - 1)]}
}

func readFiles(sourceFileName, targetFileName string) ([]byte, []byte) {
	b, err := ioutil.ReadFile(sourceFileName)
	if err != nil {
		fmt.Print(err)
	}

	a, err := ioutil.ReadFile(targetFileName) 
	if err != nil {
		fmt.Print(err)
	}

	return b, a
}

func main() {
	var sources, targets []byte = readFiles("existing-idea.txt", "targets.txt")
	idea := createIdea(sources, targets)
	fmt.Println(idea)
}

