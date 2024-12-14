package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Rules map[int]*Page
type Page struct {
	number int
	deps   Rules
}
type Updates [][]int

func parseInput() (Rules, Updates) {
	rules := make(Rules, 2048)
	reader := bufio.NewReader(os.Stdin)
	// parse rules
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if len(line) == 1 {
			break
		}
		nums := strings.Split(strings.TrimSpace(line), "|")

		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		if rules[n1] == nil {
			rules[n1] = &Page{number: n1}
		}

		n2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		if rules[n2] == nil {
			rules[n2] = &Page{number: n1}
		}

		if rules[n2].deps == nil {
			rules[n2].deps = make(Rules)
		}
		rules[n2].deps[n1] = rules[n1]
	}

	updates := make(Updates, 1024)
	eof := false
	// parse updates
	for i := 0; !eof; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				eof = true
			} else {
				panic(err)
			}
		}
		if len(line) <= 1 {
			break
		}
		nums := strings.Split(strings.TrimSpace(line), ",")

		updates[i] = make([]int, len(nums))
		for j := range len(nums) {
			n, err := strconv.Atoi(nums[j])
			if err != nil {
				panic(err)
			}
			updates[i][j] = n
		}
	}

	return rules, updates
}

func validateDeps(page *Page, present *map[int]bool, visited *map[int]bool, valid *map[int]bool) bool {
	for k, v := range page.deps {
		if !(*present)[k] || (*valid)[k] {
			continue
		}
		if !(*visited)[k] {
			return false
		}
		if !validateDeps(v, present, visited, valid) {
			return false
		}
		(*valid)[k] = true
	}
	return true
}

func partOne(rules Rules, updates Updates) int {
	sum := 0
	for i := range len(updates) {
		if updates[i] == nil {
			break
		}

		present := make(map[int]bool, len(updates[i]))
		for j := range len(updates[i]) {
			present[updates[i][j]] = true
		}
		visited := make(map[int]bool)
		valid := make(map[int]bool)

		for j := range len(updates[i]) {
			if !validateDeps(rules[updates[i][j]], &present, &visited, &valid) {
				goto nextRule
			}
			visited[updates[i][j]] = true
		}
		sum += updates[i][len(updates[i])/2]

	nextRule:
	}

	return sum
}

func moveToFront(update []int, idx int) []int {
	item := update[idx]
	for i := idx; i > 0; i-- {
		update[i] = update[i-1]
	}
	update[0] = item
	return update
}

func validateUpdate(rules Rules, update []int) bool {
	present := make(map[int]bool, len(update))
	for j := range len(update) {
		present[update[j]] = true
	}
	visited := make(map[int]bool)
	valid := make(map[int]bool)

	for i := range len(update) {
		if !validateDeps(rules[update[i]], &present, &visited, &valid) {
			return false
		}
		visited[update[i]] = true
	}
	return true
}

func fixUpdate(rules Rules, update []int) []int {
	visited := make(map[int]bool)

	for _, page := range update {
		visited[page] = true
		for dep_page := range rules[page].deps {
			if !visited[dep_page] {
				for i, p := range update {
					if dep_page == p {
						return fixUpdate(rules, moveToFront(update, i))
					}
				}
			}
		}
	}
	return update
}

func partTwo(rules Rules, updates Updates) int {
	sum := 0
	for _, update := range updates {
		if validateUpdate(rules, update) {
			continue
		}
		update = fixUpdate(rules, update)
		sum += update[len(update)/2]
	}

	return sum
}

func main() {
	rules, updates := parseInput()
	// fmt.Println(partOne(rules, updates))
	fmt.Println(partTwo(rules, updates))
}
