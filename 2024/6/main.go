package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// const InputLength = 10
const InputLength = 130

type Grid struct {
	cells [][]Cell
}

func (g *Grid) copy() *Grid {
	newGrid := Grid{
		cells: make([][]Cell, len(g.cells)),
	}
	for i, row := range g.cells {
		newGrid.cells[i] = make([]Cell, len(row))
		for j, cell := range row {
			newGrid.cells[i][j] = *cell.copy()
		}
	}
	return &newGrid
}

type Cell struct {
	visitMap map[Direction]int
	visited  int
	blocked  bool
	row      int
	col      int
}

func (c *Cell) copy() *Cell {
	newCell := Cell{
		visitMap: make(map[Direction]int, len(c.visitMap)),
		visited:  c.visited,
		blocked:  c.blocked,
		row:      c.row,
		col:      c.col,
	}
	for k, v := range c.visitMap {
		newCell.visitMap[k] = v
	}
	return &newCell
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Guard struct {
	dir Direction
	row int
	col int
}

func (g *Guard) copy() *Guard {
	return &Guard{
		dir: g.dir,
		row: g.row,
		col: g.col,
	}
}

// return false if the guard left the map
func (g *Guard) move(grid *Grid) bool {
	var nextCell *Cell
	switch g.dir {
	case Up:
		if g.row <= 0 {
			return false
		}
		nextCell = &grid.cells[g.row-1][g.col]
	case Down:
		if g.row >= len(grid.cells)-1 {
			return false
		}
		nextCell = &grid.cells[g.row+1][g.col]
	case Left:
		if g.col <= 0 {
			return false
		}
		nextCell = &grid.cells[g.row][g.col-1]
	case Right:
		if g.col >= len(grid.cells[0])-1 {
			return false
		}
		nextCell = &grid.cells[g.row][g.col+1]
	}

	if nextCell.blocked {
		g.dir += 1
		if g.dir > Left {
			g.dir = Up
		}
		return g.move(grid)
	}

	g.row = nextCell.row
	g.col = nextCell.col
	nextCell.visited = 1
	nextCell.visitMap[g.dir] += 1

	return true
}

func printGrid(grid *Grid, guard *Guard) {
	w := len(grid.cells[0])
	for i := range len(grid.cells) {
		s := make([]byte, w)
		for j, cell := range grid.cells[i] {
			if guard.row == cell.row && guard.col == cell.col {
				switch {
				case guard.dir == Up:
					s[j] = '^'
				case guard.dir == Down:
					s[j] = 'V'
				case guard.dir == Left:
					s[j] = '<'
				case guard.dir == Right:
					s[j] = '>'
				}
				continue
			}
			switch {
			case cell.blocked:
				s[j] = '#'
			case cell.visited != 0:
				s[j] = 'X'
			default:
				s[j] = '.'
			}
		}
		fmt.Println(string(s))
	}
}

func parseInput() (Grid, Guard) {
	grid := Grid{cells: make([][]Cell, InputLength)}
	guard := Guard{}

	reader := bufio.NewReader(os.Stdin)
	eof := false
	for i := 0; !eof && i < InputLength; i++ {
		row, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				eof = true
			} else {
				panic(err)
			}
		}
		row = bytes.TrimSpace(row)
		// DEBUG
		// fmt.Println(string(row))
		cells := make([]Cell, len(row))
		for j, cell := range row {
			cells[j] = Cell{visitMap: make(map[Direction]int), visited: 0, blocked: false, row: i, col: j}

			switch {
			case cell == '#':
				cells[j].blocked = true
			case cell != '.':
				cells[j].visited = 1
				guard.row = i
				guard.col = j
				switch {
				case cell == '^':
					guard.dir = Up
				case cell == '>':
					guard.dir = Right
				case cell == '<':
					guard.dir = Left
				case cell == 'V':
					guard.dir = Down
				}
				cells[j].visitMap[guard.dir] = 1
			}
		}
		grid.cells[i] = cells
	}
	return grid, guard
}

func partOne(grid Grid, guard Guard) int {
	sum := 0
	for guard.move(&grid) {
		// DEBUG
		// printGrid(&grid, &guard)
	}

	for _, row := range grid.cells {
		for _, cell := range row {
			sum += cell.visited
		}
	}

	return sum
}

func isLoop(grid *Grid, guard *Guard) int {
	for guard.move(grid) {
		// printGrid(grid, guard)
		if grid.cells[guard.row][guard.col].visitMap[guard.dir] > 1 {
			return 1
		}
	}
	return 0
}

func partTwo(grid Grid, guard Guard) int {
	sum := 0
	for i, row := range grid.cells {
		for j, cell := range row {
			if cell.blocked || cell.visited != 0 {
				continue
			}
			newGrid := grid.copy()
			newGrid.cells[i][j].blocked = true
			// DEBU
			// fmt.Println("inserted at", i, j)
			sum += isLoop(newGrid, guard.copy())
			grid.cells[i][j].blocked = false
		}
	}

	return sum
}

func main() {
	grid, guard := parseInput()
	// fmt.Println(partOne(grid, guard))
	fmt.Println(partTwo(grid, guard))
}
