package day1

import (
	"io/ioutil"
	"log"
)

// Day1 solves http://adventofcode.com/day/1
func Solve(inputFilePath string) {
	var input = readDirections(inputFilePath)
	var floor = 0
	var firstBasementPosition = 0

	for i := range input {
		switch input[i] {
		case '(':
			floor++
		case ')':
			floor--
		default:
			log.Printf("unknown input %s", input[i])
		}

		if floor == -1 && firstBasementPosition == 0 {
			firstBasementPosition = i + 1
		}
	}

	log.Printf("santa is at floor %d", floor)
	log.Printf("he entered the basement first at position %d", firstBasementPosition)
}

func readDirections(path string) (result []byte) {
	var err error
	if result, err = ioutil.ReadFile(path); err != nil {
		log.Fatal(err)
	}
	return
}
