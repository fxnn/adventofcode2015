package day4

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Solve solves http://adventofcode.com/day/4
func Solve(inputFilePath string) {
	var key = readKey(inputFilePath)
	log.Printf("Mining advent coins for key %s (len %d)", key, len(key))

	i := 1
	for {
		var hash = adventCoin(key, i)
		if strings.HasPrefix(hash, "00000") && !strings.HasPrefix(hash, "000000") {
			log.Printf("The first number is %d", i)
		}
		if strings.HasPrefix(hash, "000000") {
			log.Printf("The second number is %d", i)
			os.Exit(0)
		}
		if i%100000 == 0 {
			log.Printf("%d ...", i)
		}
		i++
	}
}

func adventCoin(key string, number int) string {
	var input = []byte(key + strconv.Itoa(number))
	var md5sum = md5.Sum(input)
	return hex.EncodeToString(md5sum[:])
}

func readKey(path string) string {
	var (
		err    error
		result []byte
	)

	if result, err = ioutil.ReadFile(path); err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(result))
}
