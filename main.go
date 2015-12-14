package main

import (
	"flag"
	"log"

	"github.com/fxnn/adventofcode2015/day1"
	"github.com/fxnn/adventofcode2015/day2"
	"github.com/fxnn/adventofcode2015/day3"
)

func main() {
	flag.Parse()

	var day = flag.Arg(0)
	var inputFilePath = flag.Arg(1)
	if inputFilePath == "" {
		inputFilePath = "day" + day + "/input.txt"
	}

	switch day {
	case "1":
		day1.Solve(inputFilePath)
	case "2":
		day2.Solve(inputFilePath)
	case "3":
		day3.Solve(inputFilePath)
	default:
		log.Printf("Day not solved: %s", day)
	}
}
