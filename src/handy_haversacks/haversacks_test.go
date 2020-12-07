package main

import "testing"

// light red bags contain 1 bright white bag, 2 muted yellow bags.
// dark orange bags contain 3 bright white bags, 4 muted yellow bags.
// bright white bags contain 1 shiny gold bag.
// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
// shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
// dark olive bags contain 3 faded blue bags, 4 dotted black bags.
// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
// faded blue bags contain no other bags.
// dotted black bags contain no other bags.

func InitRules() Rules {
	r := Rules{
		Rule: make(map[string][]Capacity),
	}

	r.Rule["light red"] = []Capacity{
		{
			Color:  "bright white",
			Number: 1,
		},
		{
			Color:  "muted yellow",
			Number: 2,
		},
	}
	r.Rule["dark orange"] = []Capacity{
		{
			Color:  "bright white",
			Number: 3,
		},
		{
			Color:  "muted yellow",
			Number: 4,
		}}

	r.Rule["bright white"] = []Capacity{
		{
			Color:  "shiny gold",
			Number: 1,
		},
	}

	r.Rule["muted yellow"] = []Capacity{
		{
			Color:  "shiny gold",
			Number: 2,
		},
		{
			Color:  "faded blue",
			Number: 9,
		},
	}
	r.Rule["shiny gold"] = []Capacity{
		{
			Color:  "dark olive",
			Number: 1,
		},
		{
			Color:  "vibrant plum",
			Number: 2,
		},
	}
	r.Rule["dark olive"] = []Capacity{
		{
			Color:  "faded blue",
			Number: 3,
		},
		{
			Color:  "dotted black",
			Number: 4,
		},
	}

	r.Rule["vibrant plum"] = []Capacity{
		{
			Color:  "faded blue",
			Number: 5,
		},
		{
			Color:  "dotted black",
			Number: 6,
		},
	}
	r.Rule["faded blue"] = []Capacity{}
	r.Rule["dotted black"] = []Capacity{}
	return r
}

func TestAreAllInvalid(t *testing.T) {
	b := InitRules()

	want := 4
	got := b.HowManyBagColorsCanContainShinyBag()

	if got != want {
		t.Errorf("got: %d, instead of: %d.", got, want)
	}

}
