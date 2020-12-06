package customs

import "fmt"

type CustomsGroups struct {
	Groups []Group
}

func (c CustomsGroups) GetSum() int {
	return 0
}

func (c *CustomsGroups) AddGroup(g Group) []Group {
	c.Groups = append(c.Groups, g)
	return c.Groups
}

type Group struct {
	Answers []Answer
}

func (g *Group) AddAnswers(a Answer) []Answer {
	g.Answers = append(g.Answers, a)
	return g.Answers
}
type Answer struct {
	questions string
}



func main() {
	questions := "abcdefghijklmnopqrstuvwxyz"
	fmt.Printf("questions: %s", questions)
}