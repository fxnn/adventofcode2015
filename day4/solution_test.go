package day4

import (
	"strings"
	"testing"
)

func TestExampleInput1(t *testing.T) {

	var actual = adventCoin("abcdef", 609043)
	if !strings.HasPrefix(actual, "000001dbbfa") {
		t.Errorf("Code should start with 000001dbbfa, but was %s", actual)
	}

}

func TestExampleInput2(t *testing.T) {

	var actual = adventCoin("pqrstuv", 1048970)
	if !strings.HasPrefix(actual, "000006136ef") {
		t.Errorf("Code should start with 000006136ef, but was %s", actual)
	}

}
