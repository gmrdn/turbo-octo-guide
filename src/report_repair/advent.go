package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ExpenseReport struct {
	Entries []ExpenseEntry
}

type ExpenseEntry struct {
	value int
}

func (report *ExpenseReport) AddExpense(entry ExpenseEntry) []ExpenseEntry {
	report.Entries = append(report.Entries, entry)
	return report.Entries
}

func (report *ExpenseReport) findPairMaking(targetValue int) [2]int {
	for i := 0; i < len(report.Entries); i++ {
		indexFirstNumber := i
		valueFirstNumber := report.Entries[i].value
		for j := 0; j < len(report.Entries); j++ {
			indexSecondNumber := j
			valueSecondNumber := report.Entries[j].value
			if indexSecondNumber != indexFirstNumber && valueFirstNumber+valueSecondNumber == targetValue {
				return [2]int{valueFirstNumber, valueSecondNumber}
			}
		}
	}
	return [2]int{0, 0}
}

func (report *ExpenseReport) FindPairMaking2020() [2]int {
	return report.findPairMaking(2020)
}

func (report *ExpenseReport) FindThreeMaking2020() [3]int {
	for i := 0; i < len(report.Entries); i++ {
		valueFirstNumber := report.Entries[i].value
		target := 2020 - valueFirstNumber
		twoOthers := report.findPairMaking(target)
		if twoOthers[0] != 0 && twoOthers[1] != 0 {
			return [3]int{valueFirstNumber, twoOthers[0], twoOthers[1]}
		}
	}
	return [3]int{0, 0, 0}
}

func main() {

	fileHandle, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	report := ExpenseReport{
		Entries: []ExpenseEntry{},
	}

	for fileScanner.Scan() {
		if s, err := strconv.Atoi(fileScanner.Text()); err == nil {
			report.AddExpense(ExpenseEntry{s})
		}
	}

	pair := report.FindPairMaking2020()

	fmt.Printf("First value: %d, second value: %d, product: %d\n", pair[0], pair[1], pair[0]*pair[1])

	tripplet := report.FindThreeMaking2020()

	fmt.Printf("First value: %d, second value: %d, third value: %d,  product: %d\n", tripplet[0], tripplet[1], tripplet[2], tripplet[0]*tripplet[1]*tripplet[2])

}
