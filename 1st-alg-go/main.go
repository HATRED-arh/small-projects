package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var words []string
var indexes []int

func main() {

	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	sort.Strings(words)
	for i := 0; i < len(words)-1; i++ {
		if words[i] != words[i+1] {
			indexes = append(indexes, i+1)
		}
	}

	indexes = append(indexes, 0)
	copy(indexes[1:], indexes)
	indexes[0] = 0
	indexes = append(indexes, len(words))
	for i := 0; i < len(indexes)-1; i++ {
		fmt.Printf("Word \"%s\" repeats %d %s \n", words[indexes[i]], indexes[i+1]-indexes[i], corrector(indexes[i+1]-indexes[i]))
	}
}

func corrector(quantity int) string {
	if quantity == 1 {
		return "time"
	} else {
		return "times"
	}
}
