package counter

import "strings"

// SplitLines receive an array of strings and split each record and produce words to an output channel.
func SplitLines(lines []string) <-chan string {
	output := make(chan string, 5)
	go func() {
		defer close(output)
		for _, line := range lines {
			for _, word := range strings.Split(line, " ") {
				output <- word
			}
		}
	}()
	return output
}

func WordNormalization(input <-chan string) <-chan string {
	output := make(chan string, 5)
	go func() {
		defer close(output)
		for word := range input {
			output <- strings.ToLower(word)
		}
	}()
	return output
}

// CountWord ..
func CountWord(input <-chan string) map[string]int {
	wordCountSummary := make(map[string]int)
	for word := range input {
		if _, ok := wordCountSummary[word]; ok {
			wordCountSummary[word]++
		} else {
			wordCountSummary[word] = 1
		}
	}
	return wordCountSummary
}
