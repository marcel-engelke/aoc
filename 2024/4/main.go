package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	InputLength = 140
	// InputLength = 10
)

func readInput() [][]byte {
	data := make([][]byte, InputLength)
	reader := bufio.NewReader(os.Stdin)
	for i := range InputLength {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		data[i] = line
	}
	return data
}

type Vec [4][2]int

func partOne(data [][]byte) int {
	check := func(vec Vec) int {
		if data[vec[0][0]][vec[0][1]] == 'X' &&
			data[vec[1][0]][vec[1][1]] == 'M' &&
			data[vec[2][0]][vec[2][1]] == 'A' &&
			data[vec[3][0]][vec[3][1]] == 'S' {
			return 1
		}

		return 0
	}

	sum := 0
	for i := range len(data) {
		for j := range len(data[i]) {
			// right
			if j <= len(data[i])-4 {
				sum += check(Vec{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}})
			}
			// left
			if j >= 3 {
				sum += check(Vec{{i, j}, {i, j - 1}, {i, j - 2}, {i, j - 3}})
			}
			// up
			if i >= 3 {
				sum += check(Vec{{i, j}, {i - 1, j}, {i - 2, j}, {i - 3, j}})
			}
			// down
			if i <= len(data)-4 {
				sum += check(Vec{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}})
			}
			// left diag up
			if i >= 3 && j >= 3 {
				sum += check(Vec{{i, j}, {i - 1, j - 1}, {i - 2, j - 2}, {i - 3, j - 3}})
			}
			// // left diag down
			if i <= len(data)-4 && j >= 3 {
				sum += check(Vec{{i, j}, {i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}})
			}
			// right diag up
			if i >= 3 && j <= len(data[i])-4 {
				sum += check(Vec{{i, j}, {i - 1, j + 1}, {i - 2, j + 2}, {i - 3, j + 3}})
			}
			// right diag down
			if i <= len(data)-4 && j <= len(data[i])-4 {
				sum += check(Vec{{i, j}, {i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}})
			}

		}
	}
	return sum
}

// MMASS
type Vec2 [5][2]int

func partTwo(data [][]byte) int {
	check := func(vec Vec2) int {
		if data[vec[0][0]][vec[0][1]] == 'M' &&
			data[vec[1][0]][vec[1][1]] == 'M' &&
			data[vec[2][0]][vec[2][1]] == 'A' &&
			data[vec[3][0]][vec[3][1]] == 'S' &&
			data[vec[4][0]][vec[4][1]] == 'S' {
			return 1
		}

		return 0
	}

	//   01234 j
	// 0 MMMSX
	// 1 MSAMS
	// 2 AMXSA
	// 3 MSAMX
	// 4 XMASM
	// i

	sum := 0
	for i := range len(data) {
		for j := range len(data[i]) {
			// left
			if i <= len(data)-3 && j <= len(data[i])-3 {
				sum += check(Vec2{{i, j}, {i + 2, j}, {i + 1, j + 1}, {i, j + 2}, {i + 2, j + 2}})
			}
			// right
			if i <= len(data)-3 && j >= 2 {
				sum += check(Vec2{{i, j}, {i + 2, j}, {i + 1, j - 1}, {i, j - 2}, {i + 2, j - 2}})
			}
			// up
			if i >= 2 && j <= len(data[i])-3 {
				sum += check(Vec2{{i, j}, {i, j + 2}, {i - 1, j + 1}, {i - 2, j}, {i - 2, j + 2}})
			}
			// down
			if i <= len(data)-3 && j <= len(data[i])-3 {
				sum += check(Vec2{{i, j}, {i, j + 2}, {i + 1, j + 1}, {i + 2, j}, {i + 2, j + 2}})
			}
		}
	}
	return sum
}

func main() {
	data := readInput()
	// fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}
