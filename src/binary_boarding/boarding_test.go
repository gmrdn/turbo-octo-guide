package main

import "testing"


func InitPasses() BoardingPasses {
	b := BoardingPasses{
		Seats: []Seat{},
	}
	b.AddSeat(Seat{
		Specification: "FBFBBFFRLR",
		SeatID: 0,
	})
	b.AddSeat(Seat{
		Specification: "BFFFBBFRRR",
		SeatID: 0,
	})
	b.AddSeat(Seat{
		Specification: "FFFBBBFRRR",
		SeatID: 0,
	})
	b.AddSeat(Seat{
		Specification: "BBFFBBFRLL",
		SeatID: 0,
	})
	return b
}


func TestCalculateSeatID(t *testing.T) {
	p := Plane{127, 7}
	b := InitPasses()
	b.Seats[0].CalcAndUpdateID(p)

	want := 357
	got := b.Seats[0].SeatID
	if got != want {
		t.Errorf("got: %d, instead of: %d.", got, want)
	}
}