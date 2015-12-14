package day3

import "testing"

func TestPart1_Example1(t *testing.T) {

	var s = newSolution()
	s.followDirections("^v", 1)

	if s.countHousesWithPresents != 2 {
		t.Errorf("Expected 2 houses with presents, but got %d", s.countHousesWithPresents)
	}

}

func TestPart1_Example2(t *testing.T) {

	var s = newSolution()
	s.followDirections("^>v<", 1)

	if s.countHousesWithPresents != 4 {
		t.Errorf("Expected 4 houses with presents, but got %d", s.countHousesWithPresents)
	}

}

func TestPart1_Example3(t *testing.T) {

	var s = newSolution()
	s.followDirections("^v^v^v^v^v", 1)

	if s.countHousesWithPresents != 2 {
		t.Errorf("Expected 2 houses with presents, but got %d", s.countHousesWithPresents)
	}

}

func TestPart2_Example1(t *testing.T) {

	var s = newSolution()
	s.followDirections("^v", 2)

	if s.countHousesWithPresents != 3 {
		t.Errorf("Expected 3 houses with presents, but got %d", s.countHousesWithPresents)
	}

}

func TestPart2_Example2(t *testing.T) {

	var s = newSolution()
	s.followDirections("^>v<", 2)

	if s.countHousesWithPresents != 3 {
		t.Errorf("Expected 3 houses with presents, but got %d", s.countHousesWithPresents)
	}

}

func TestPart2_Example3(t *testing.T) {

	var s = newSolution()
	s.followDirections("^v^v^v^v^v", 2)

	if s.countHousesWithPresents != 11 {
		t.Errorf("Expected 11 houses with presents, but got %d", s.countHousesWithPresents)
	}

}
