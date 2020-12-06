package main

import "testing"



func InitGroups() CustomsGroups {
	c := CustomsGroups{
		Groups: []Group{},
	}
	c.AddGroup(Group{
		Answers: []Answer{},
	})
	c.AddGroup(Group{
		Answers: []Answer{},
	})
	c.AddGroup(Group{
		Answers: []Answer{},
	})
	c.AddGroup(Group{
		Answers: []Answer{},
	})
	c.AddGroup(Group{
		Answers: []Answer{},
	})
	

	c.Groups[0].AddAnswers(Answer{
		questions: "abc",
	})
	c.Groups[1].AddAnswers(Answer{
		questions: "a",
	})
	c.Groups[1].AddAnswers(Answer{
		questions: "b",
	})
	c.Groups[1].AddAnswers(Answer{
		questions: "c",
	})
	c.Groups[2].AddAnswers(Answer{
		questions: "ab",
	})
	c.Groups[2].AddAnswers(Answer{
		questions: "ac",
	})
	c.Groups[3].AddAnswers(Answer{
		questions: "a",
	})
	c.Groups[3].AddAnswers(Answer{
		questions: "a",
	})
	c.Groups[3].AddAnswers(Answer{
		questions: "a",
	})
	c.Groups[3].AddAnswers(Answer{
		questions: "a",
	})
	c.Groups[4].AddAnswers(Answer{
		questions: "b",
	})

	return c
}

func TestGroups(t *testing.T) {
	c := InitGroups()
	want := 11
	got := c.GetSum()
	if got != want {
		t.Errorf("got: %d, instead of: %d ", got, want)
	}
}