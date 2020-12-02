package main

import "testing"

func InitExpenseReport() ExpenseReport {
	report := ExpenseReport{
		Entries: []ExpenseEntry{},
	}
	report.AddExpense(ExpenseEntry{1721})
	report.AddExpense(ExpenseEntry{979})
	report.AddExpense(ExpenseEntry{366})
	report.AddExpense(ExpenseEntry{299})
	report.AddExpense(ExpenseEntry{675})
	report.AddExpense(ExpenseEntry{1456})

	return report
}

func TestFindPairMaking2020(t *testing.T) {

	report := InitExpenseReport()

	want := [2]int{1721, 299}
	got := report.FindPairMaking2020()
	if got != want {
		t.Errorf("got: %d, instead of: %d.", got, want)
	}
}

func TestFindThreeMaking2020(t *testing.T) {

	report := InitExpenseReport()

	want := [3]int{979, 366, 675}
	got := report.FindThreeMaking2020()
	if got != want {
		t.Errorf("got: %d, instead of: %d.", got, want)
	}
}
