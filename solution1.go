package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var input = read(inputFilePath())
	var floor = 0
	for i := range input {
		switch input[i] {
		case '(':
			floor++
		case ')':
			floor--
		default:
			log.Printf("unknown input %s", input[i])
		}
	}
	log.Printf("santa is at floor %d", floor)
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
