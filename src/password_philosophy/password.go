package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PasswordList struct {
	Entries []PasswordEntry
}
type PasswordEntry struct {
	minimum  int
	maximum  int
	letter   string
	password string
}

func (pwd *PasswordEntry) IsValid() bool {
	nb := strings.Count(pwd.password, pwd.letter)
	if nb >= pwd.minimum && nb <= pwd.maximum {
		return true
	}
	return false
}

func (list *PasswordList) AddPassword(s string) []PasswordEntry {
	split := strings.Split(s, " ")
	numbers := strings.Split(split[0], "-")
	minimum, _ := strconv.Atoi(numbers[0])
	maximum, _ := strconv.Atoi(numbers[1])
	letter := strings.Split(split[1], ":")[0]
	password := split[2]
	entry := PasswordEntry{
		minimum:  minimum,
		maximum:  maximum,
		password: password,
		letter:   letter,
	}
	list.Entries = append(list.Entries, entry)
	return list.Entries
}

func (list *PasswordList) GetNbValid() int {
	nb := 0
	for i := 0; i < len(list.Entries); i++ {
		if list.Entries[i].IsValid() {
			nb++
		}
	}
	return nb
}

func main() {
	fmt.Printf("hello")
	fileHandle, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	passwords := PasswordList{
		Entries: []PasswordEntry{},
	}

	for fileScanner.Scan() {
		s := fileScanner.Text()
		passwords.AddPassword(s)
	}

	nb := passwords.GetNbValid()
	fmt.Printf("Number of valid passwords: %d\n", nb)
}
