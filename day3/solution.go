package day3

import (
	"io/ioutil"
	"log"
)

// Solve solves http://adventofcode.com/day/3
func Solve(inputFilePath string) {
	var directions = readDirections(inputFilePath)

	var part1 = newSolution()
	part1.followDirections(directions, 1)
	log.Printf("In year one, %d houses receive at least one present", part1.countHousesWithPresents)

	var part2 = newSolution()
	part2.followDirections(directions, 2)
	log.Printf("In year two, %d houses receive at least one present", part2.countHousesWithPresents)
}

// point defines the coordinates of a house
type point struct {
	x int
	y int
}

func (p *point) move(d rune) {
	switch d {
	case '<':
		p.x -= 1
	case '>':
		p.x += 1
	case '^':
		p.y += 1
	case 'v':
		p.y -= 1
	default:
		log.Fatalf("Unknown direction: %v", d)
	}
}

// house represents the number of presents it became
type house int

type solution struct {
	countHousesWithPresents int
	houses                  map[point]house
}

func newSolution() solution {
	return solution{0, make(map[point]house)}
}

func (s *solution) visit(pos *point) {
	if s.houses[*pos] == 0 {
		s.countHousesWithPresents++
	}
	s.houses[*pos]++
}

func (s *solution) followDirections(directions string, countSantas int) {
	var santas = make([]*point, countSantas)

	for i := range santas {
		santas[i] = &point{0, 0}
		s.visit(santas[i])
	}

	for i, direction := range directions {
		var santa = santas[i%countSantas]
		santa.move(direction)
		s.visit(santa)
	}

}

func readDirections(path string) string {
	var (
		err    error
		result []byte
	)

	if result, err = ioutil.ReadFile(path); err != nil {
		log.Fatal(err)
	}

	return string(result)
}
