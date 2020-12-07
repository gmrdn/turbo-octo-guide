package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rules struct {
	Rule map[string][]Capacity
}

func (r *Rules) HowManyBagColorsCanContainShinyBag() int {
	nbContainer := 0
	for key, _ := range r.Rule {
		if r.CanContain(key, "shiny gold") {
			nbContainer++
		}

	}

	return nbContainer
}

func (r *Rules) CanContain(key string, find string) bool {
	rule := r.Rule[key]
	for _, c := range rule {
		if c.Color == find {
			return true
		} else {
			return r.CanContain(c.Color, find)
		}

	}
	return false
}

type Capacity struct {
	Color  string
	Number int
}

func main() {
	r := Rules{
		Rule: make(map[string][]Capacity),
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

		contained := strings.Split(s, "contain ")[1]
		contained = contained[:len(contained)-1] // enlève le point
		contained = strings.Replace(contained, "bags", "", -1)
		contained = strings.Replace(contained, "bag", "", -1)
		if contained == "no other" {
			r.Rule[container] = []Capacity{}
		} else {
			bags := strings.Split(contained, ", ")
			capacities := []Capacity{}
			for _, b := range bags {
				capArray := strings.Split(b, " ")
				nb, _ := strconv.Atoi(capArray[0])
				color := capArray[1] + " " + capArray[2]
				capacities = append(capacities, Capacity{
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
