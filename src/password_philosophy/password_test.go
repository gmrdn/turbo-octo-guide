package main

import "testing"

func PrepareList() PasswordList {
	list := PasswordList{
		Entries: []PasswordEntry{},
	}
	list.AddPassword("1-3 a: abcde")
	list.AddPassword("1-3 b: cdefg")
	list.AddPassword("2-9 c: ccccccccc")
	return list
}

func TestGetNbValid(t *testing.T) {
	list := PrepareList()
	want := 1
	got := list.GetNbValid()
	if got != want {
		t.Errorf("got: %d, instead of: %d.", got, want)
	}
}

func TestIsValid(t *testing.T) {
	password := PasswordEntry{
		minimum:  1,
		maximum:  3,
		letter:   "a",
		password: "abcde",
	}
	want := true
	got := password.IsValid()
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}

func TestIsNotValid(t *testing.T) {
	password := PasswordEntry{
		minimum:  1,
		maximum:  3,
		letter:   "b",
		password: "cdefg",
	}
	want := false
	got := password.IsValid()
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}

func TestIsNotValidEither(t *testing.T) {
	password := PasswordEntry{
		minimum:  2,
		maximum:  9,
		letter:   "c",
		password: "ccccccccc",
	}
	want := false
	got := password.IsValid()
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}
