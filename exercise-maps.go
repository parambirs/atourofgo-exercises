package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		count, present := m[word]
		if present {
			m[word] = count + 1
		} else {
			m[word] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}

/*
> go get golang.org/x/tour/wc
> go run exercise-maps.go 
PASS
 f("I am learning Go!") =
  map[string]int{"I":1, "am":1, "learning":1, "Go!":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"The":1, "brown":1, "fox":1, "jumped":1, "over":1, "dog.":1, "quick":1, "the":1, "lazy":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"a":1, "donut.":2, "Then":1, "another":1, "I":2, "ate":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"plan":1, "canal":1, "panama.":1, "A":1, "man":1, "a":2}
*/