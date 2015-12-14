package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Day2()
}

// Day2 solves http://adventofcode.com/day/2
func Day2() {
	var presents = readPresents(inputFilePath())
	
	var area = 0
	var length = 0
	
	for i := range presents {
		// NOTE, that the dimensions are sorted in increasing order
		var dimensions = presents[i]

		area += 2 * dimensions[0] * dimensions[1]
		area += 2 * dimensions[0] * dimensions[2]
		area += 2 * dimensions[1] * dimensions[2]
		area += dimensions[0] * dimensions[1]

		length += 2*dimensions[0] + 2*dimensions[1]
		length += dimensions[0] * dimensions[1] * dimensions[2]
	}
	
	log.Printf("The elves need %d sq feet of paper", area)
	log.Printf("They also require %d feet of ribbon", length)
}

type present []int

func readPresents(path string) []present {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]present, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, parsePresent(scanner.Text()))
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}
	file.Close()

	return result
}

func inputFilePath() string {
	flag.Parse()
	return flag.Arg(0)
}

func parsePresent(input string) present {
	var result = make(present, 3)
	var parts = strings.Split(input, "x")
	if len(parts) != 3 {
		log.Fatalf("line without 3 dimensions: %s", input)
	}
	for i := range parts {
		val, err := strconv.Atoi(parts[i])
		if err != nil {
			log.Fatal(err)
		}
		result[i] = val
	}
	sort.Ints(result) // NOTE, that we sort the dimensions
	return result
}
