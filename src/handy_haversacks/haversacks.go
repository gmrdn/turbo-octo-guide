package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//https://fodor.org/blog/go-graph/
type Graph struct {
	Nodes []*GraphNode
}

type GraphNode struct {
	Label string
	Edges map[string]string
}

func New() *Graph {
	return &Graph{
		Nodes: []*GraphNode{},
	}
}

type Rules struct {
	Rule map[string][]Rule
}

func (r *Rules) HowManyBagColorsCanContainShinyBag() int {
	nbContainer := 0

	for key, _ := range r.Rule {
		for _, c := range r.Rule[key] {
			if r.CanContain(c.Color, "shiny gold") {
				nbContainer++
			}
		}

	}

	return nbContainer
}

func (r *Rules) CanContain(key string, find string) bool {
	return false
}

type Rule struct {
	Color  string
	Number int
}

func main() {
	r := Rules{
		Rule: make(map[string][]Rule),
	}

	fileHandle, err := os.Open("day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {

		s := fileScanner.Text()
		container := strings.Split(s, "contain ")[0]
		container = container[:len(container)-6] // enlève " bags "
		fmt.Printf("Container: %s\n", container)

		contained := strings.Split(s, "contain ")[1]
		contained = contained[:len(contained)-1] // enlève le point
		contained = strings.Replace(contained, "bags", "", -1)
		contained = strings.Replace(contained, "bag", "", -1)
		if contained == "no other" {
			r.Rule[container] = []Rule{}
			fmt.Printf("Empty capacity added\n")

		} else {
			bags := strings.Split(contained, ", ")
			capacities := []Rule{}

			for _, b := range bags {
				capArray := strings.Split(b, " ")
				nb, _ := strconv.Atoi(capArray[0])
				color := capArray[1] + " " + capArray[2]

				capacities = append(capacities, Rule{
					Color:  color,
					Number: nb,
				})

				fmt.Printf("Capacity added: %s, %d\n", color, nb)
			}
			r.Rule[container] = capacities
		}

	}

	fmt.Printf("Nb %d\n", r.HowManyBagColorsCanContainShinyBag())
}
