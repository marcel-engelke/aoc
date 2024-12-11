package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func readInput() []byte {
	reader := bufio.NewReader(os.Stdin)
	input, err := io.ReadAll(reader)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return input
}

func partOne(input []byte) int {
	re := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	matches := re.FindAll(input, -1)

	sum := 0
	re = regexp.MustCompile("\\d+")
	for i := range len(matches) {
		nums := re.FindAll(matches[i], -1)
		n1, err := strconv.Atoi(string(nums[0]))
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(string(nums[1]))
		if err != nil {
			panic(err)
		}
		sum += n1 * n2
	}
	return sum
}

func partTwo(input []byte) int {
	re := regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
	matches := re.FindAll(input, -1)

	enabled := true
	sum := 0
	re = regexp.MustCompile("\\d+")
	for i := range len(matches) {
		if matches[i][0] == 'd' {
			if matches[i][2] == '(' {
				enabled = true
			} else {
				enabled = false
			}
			continue
		}

		if !enabled {
			continue
		}

		nums := re.FindAll(matches[i], -1)
		n1, err := strconv.Atoi(string(nums[0]))
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(string(nums[1]))
		if err != nil {
			panic(err)
		}

		sum += n1 * n2
	}
	return sum
}

func main() {
	input := readInput()
	// fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
