package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var input = read(inputFilePath())
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

func read(path string) (result []byte) {
	var err error
	if result, err = ioutil.ReadFile(path); err != nil {
		log.Fatal(err)
	}
	return
}

func inputFilePath() string {
	flag.Parse()
	return flag.Arg(0)
}
