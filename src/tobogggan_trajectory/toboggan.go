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

		if currentPosition[1] < area.Height {
			if area.IsTree(currentPosition[0], currentPosition[1]) {
				nb++
			}
		}
	}
	return nb
}

type Row struct {
	Trees []int
}

func main() {
	fmt.Printf("hello")

	area := Area{
		Rows:   []Row{},
		Width:  0,
		Height: 0,
	}

	fileHandle, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	nbrows := 0
	for fileScanner.Scan() {
		s := fileScanner.Text()
		area.Width = len(s)
		area.AddRow(s)
		nbrows++
	}
	area.Height = nbrows

	slope1 := Slope{x: 1, y: 1}
	nb1 := area.HowManyTrees(slope1)
	fmt.Printf("Number of trees: %d\n", nb1)
	slope2 := Slope{x: 3, y: 1}
	nb2 := area.HowManyTrees(slope2)
	fmt.Printf("Number of trees: %d\n", nb2)
	slope3 := Slope{x: 5, y: 1}
	nb3 := area.HowManyTrees(slope3)
	fmt.Printf("Number of trees: %d\n", nb3)
	slope4 := Slope{x: 7, y: 1}
	nb4 := area.HowManyTrees(slope4)
	fmt.Printf("Number of trees: %d\n", nb4)
	slope5 := Slope{x: 1, y: 2}
	nb5 := area.HowManyTrees(slope5)
	fmt.Printf("Number of trees: %d\n", nb5)

	fmt.Printf("Total: %d\n", nb1*nb2*nb3*nb4*nb5)

}
