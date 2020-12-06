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
	b.Seats[0].CalcAndLocateSeat(p)

	want := 357
	got := b.Seats[0].SeatID
	if got != want {
		t.Errorf("got: %d, instead of: %d.", got, want)
	}
}

func TestFindEmptySeats(t *testing.T) {
	p := Plane{2, 0}
	b := BoardingPasses{
		Seats: []Seat{}}

	b.AddSeat(Seat{
		Specification: "FFL",
	})
	b.AddSeat(Seat{
		Specification: "BL",
	})

	b.Seats[0].CalcAndLocateSeat(p)
	b.Seats[1].CalcAndLocateSeat(p)

	want := []Seat{Seat{
		Specification: "",
		Row: 1,
		Col: 0,
	}}
	got := p.FindEmptySeats(b)
	if got[0].Row != want[0].Row {
		t.Errorf("got: %d, instead of: %d.", got[0].Row, want[0].Row)
	}

}

