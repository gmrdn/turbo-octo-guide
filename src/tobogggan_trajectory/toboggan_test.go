package main

import "testing"

func PrepareArea() Area {
	testArea := Area{
		Rows:   []Row{},
		Width:  0,
		Height: 0,
	}

	testArea.AddRow("..##.......")
	testArea.AddRow("#...#...#..")
	testArea.AddRow(".#....#..#.")
	testArea.AddRow("..#.#...#.#")
	testArea.AddRow(".#...##..#.")
	testArea.AddRow("..#.##.....")
	testArea.AddRow(".#.#.#....#")
	testArea.AddRow(".#........#")
	testArea.AddRow("#.##...#...")
	testArea.AddRow("#...##....#")
	testArea.AddRow(".#..#...#.#")
	return testArea

}

func Test_3_1_IsNotTree(t *testing.T) {
	testArea := PrepareArea()
	testArea.Height = 11
	testArea.Width = 11

	want := false
	got := testArea.IsTree(3, 1)
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}

func Test_6_2_IsTree(t *testing.T) {
	testArea := PrepareArea()
	testArea.Height = 11
	testArea.Width = 11

	want := true
	got := testArea.IsTree(6, 2)
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}

func Test_12_4_IsTree(t *testing.T) {
	testArea := PrepareArea()
	testArea.Height = 11
	testArea.Width = 11

	want := true
	got := testArea.IsTree(12, 4)
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}

func TestShouldEncounter7Trees(t *testing.T) {
	testArea := PrepareArea()
	testArea.Height = 11
	testArea.Width = 11

	slope := Slope{x: 3, y: 1}

	want := 7
	got := testArea.HowManyTrees(slope)
	if got != want {
		t.Errorf("got: %d, instead of: %d.", got, want)
	}
}
