package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	InputLength = 1000
)

func parseInput() [][]int {
	reports := make([][]int, InputLength)
	reader := bufio.NewReader(os.Stdin)
	for i := range InputLength {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		cols := strings.Split(line, " ")
		for j := range len(cols) {
			col, err := strconv.Atoi(cols[j])
			if err != nil {
				panic(err)
			}
			reports[i] = append(reports[i], col)
		}
	}

	return reports
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isSafe(report []int) bool {
	decr := report[0] > report[1]
	for i := 1; i < len(report); i++ {
		diff := abs(report[i-1] - report[i])
		if diff < 1 || diff > 3 {
			return false
		}
		if decr && report[i-1] < report[i] {
			return false
		}
		if !decr && report[i] < report[i-1] {
			return false
		}
	}
	return true
}

func partOne(reports [][]int) int {
	sum := 0
	for i := range InputLength {
		if isSafe(reports[i]) {
			sum += 1
		}
	}
	return sum
}

func isSafeAdv(report []int, retry bool) bool {
	// because we now accept a bad value, determining the order is no longer as trivial
	incrs, decrs := 0, 0
	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			incrs += 1
		} else {
			decrs += 1
		}
	}
	decr := decrs > incrs

	valid := true
	i := 1
	for ; i < len(report); i++ {
		diff := abs(report[i-1] - report[i])
		if diff < 1 || diff > 3 {
			goto invalid
		}
		if decr && report[i-1] < report[i] {
			goto invalid
		}
		if !decr && report[i] < report[i-1] {
			goto invalid
		}
		continue
	invalid:
		valid = false
		break
	}

	if !valid && retry {
		// if the first comparison failed, try swapping out the first and the second values
		if i == 1 && isSafeAdv(slices.Delete(slices.Clone(report), i-1, i), false) {
			return true
		}
		return isSafeAdv(slices.Delete(report, i, i+1), false)
	}

	return valid
}

func partTwo(reports [][]int) int {
	sum := 0
	for i := range InputLength {
		if isSafeAdv(reports[i], true) {
			sum += 1
		}
	}
	return sum
}

func main() {
	reports := parseInput()
	// fmt.Println(partOne(reports))
	fmt.Println(partTwo(reports))
}
