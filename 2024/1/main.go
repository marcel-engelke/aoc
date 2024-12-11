package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	left  []int
	right []int
}

const (
	InputLength = 1000
)

func parseInput() Data {
	data := Data{left: make([]int, InputLength), right: make([]int, InputLength)}
	reader := bufio.NewReader(os.Stdin)
	for i := range InputLength {
		line, _ := reader.ReadString('\n')
		cols := strings.SplitN(line, " ", 2)

		value, err := strconv.Atoi(strings.TrimSpace(cols[0]))
		if err != nil {
			panic(err)
		}
		data.left[i] = value
		value, err = strconv.Atoi(strings.TrimSpace(cols[1]))
		if err != nil {
			panic(err)
		}
		data.right[i] = value
	}

	return data
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func partOne(data Data) int {
	sort.Slice(data.left, func(a, b int) bool {
		return data.left[a] < data.left[b]
	})
	sort.Slice(data.right, func(a, b int) bool {
		return data.right[a] < data.right[b]
	})

	sum := 0
	for i := 0; i < InputLength; i++ {
		sum += abs(data.left[i] - data.right[i])
	}
	return sum
}

func partTwo(data Data) int {
	occurences := make(map[int]int, InputLength)
	for i := range len(data.right) {
		occurences[data.right[i]] += 1
	}

	sum := 0
	for i := range InputLength {
		sum += data.left[i] * occurences[data.left[i]]
	}
	return sum
}

func main() {
	data := parseInput()

	// fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}
