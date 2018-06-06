package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	rand.Seed(time.Now().Unix())
	var sources, targets []byte = readFiles("existing-idea.txt", "targets.txt")
	idea := createIdea(sources, targets)
	fmt.Println(idea)
}

