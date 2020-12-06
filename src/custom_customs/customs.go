package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CustomsGroups struct {
	Groups []Group
}

func (c CustomsGroups) GetSum() int {
	sum := 0
	for _, g := range c.Groups {
		sum += g.GetNbAllAnswered()
	}
	return sum
}

func (c *CustomsGroups) AddGroup(g Group) []Group {
	c.Groups = append(c.Groups, g)
	return c.Groups
}

type Group struct {
	Answers []Answer
}

func (g *Group) GetNbUniqueAnswer() int {
	set := make(map[rune]struct{})

	for _, a := range g.Answers { // aaa, abc, bc, c
		for _, r := range a.questions { // a a a
			if _, ok := set[r]; ok {
			} else {
				set[r] = struct{}{}
			}
		}
	}
	return len(set)
}

func (g *Group) GetNbAllAnswered() int {
	set := make(map[rune]struct{})
	
	// first answer
	for _, r := range g.Answers[0].questions {
		if _, ok := set[r]; ok {
		} else {
			set[r] = struct{}{}
		}
	} 
	// set a b c
	for _, a := range g.Answers { // abc, ac, ad
		for r := range set {
			if strings.ContainsRune(a.questions, r) {

			} else {
				delete(set, r)
			}
		}
	}
	return len(set)
}


func (g *Group) AddAnswers(a Answer) []Answer {
	g.Answers = append(g.Answers, a)
	return g.Answers
}
type Answer struct {
	questions string
}



func main() {

	fmt.Printf("hello")

	fileHandle, err := os.Open("day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)


	c := CustomsGroups{
		Groups: []Group{},
	}

	c.AddGroup(Group{
		Answers: []Answer{},
	})
	currentIndex := 0

	for fileScanner.Scan() {
		s := fileScanner.Text()

		if s == "" {
			c.AddGroup(Group{
				Answers: []Answer{},
			})
			currentIndex++
		} else {
			c.Groups[currentIndex].AddAnswers(Answer{questions: s})
		}
	}

	fmt.Printf("Total: %d\n", c.GetSum())

}