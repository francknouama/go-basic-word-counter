package main

import (
	"fmt"

	"github.com/francknouama/go-basic-word-counter/counter"
)

func main() {

	lines := []string{
		"Peggy and I wandered back down Fifth Avenue with the rest of the crowd dribbling out of",
		"the Robert Palmer concert that had just reached its exhausted finale in Central Park. It was",
		"part of the annual Dr Pepper Central Park Music Festival and whatever Robert Palmer may",
		"have thought, I, for one, was extremely grateful for their sponsorship, because it was one of",
	}

	wordsCh := counter.SplitLines(lines)
	normWordsCh := counter.WordNormalization(wordsCh)
	summary := counter.CountWord(normWordsCh)

	for key, value := range summary {
		fmt.Printf("%s: %d\n", key, value)
	}
}
