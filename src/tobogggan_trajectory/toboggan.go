package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Slope struct {
	x int
	y int
}

type Area struct {
	Rows   []Row
	Width  int
	Height int
}

func (area *Area) AddRow(s string) []Row {
	trees := []int{}

	for pos, char := range s {
		if char == '#' {
			trees = append(trees, pos)
		}
	}
	row := Row{Trees: trees}
	area.Rows = append(area.Rows, row)
	return area.Rows
}

func (area *Area) IsTree(x int, y int) bool {
	row := area.Rows[y]
	trees := row.Trees

	for _, n := range trees {
		if x%area.Width == n {
			return true
		}
	}
	return false
}

func (area *Area) HowManyTrees(slope Slope) int {
	nb := 0
	for i := 0; i < len(area.Rows); i++ {
		currentPosition := [2]int{i * slope.x, i * slope.y}
		if area.IsTree(currentPosition[0], currentPosition[1]) {
			nb++
		}
	}
	return nb
}

type Row struct {
	Trees []int
}

func main() {
	fmt.Printf("hello")
	nb := 0
	area := Area{
		Rows:  []Row{},
		Width: 0,
	}

	fileHandle, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		s := fileScanner.Text()
		area.Width = len(s)
		area.AddRow(s)
	}

	slope := Slope{x: 3, y: 1}

	nb = area.HowManyTrees(slope)

	fmt.Printf("Number of trees: %d\n", nb)
}
