package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type BoardingPasses struct {
	Seats []Seat
}

func (b *BoardingPasses) AddSeat(s Seat) []Seat {
	b.Seats = append(b.Seats, s)
	return b.Seats
}
type Seat struct {
	Specification string
	SeatID	int
}

func (s *Seat) CalcAndUpdateID(p Plane) {
	startingRow := 0
	endingRow := p.Rows
	startingCol := 0
	endingCol := p.Columns

	for _, char := range s.Specification {
		middleRow := (startingRow + endingRow) / 2
		middleCol := (startingCol + endingCol) / 2
		if char == 'F' {
			endingRow = middleRow
		}
		if char == 'B' {
			startingRow = middleRow + 1
		}
		if char == 'R' {
			startingCol = middleCol + 1
		}
		if char == 'L' {
			endingCol = middleCol
		}
	}
	
	s.SeatID = endingRow * 8 + endingCol
}


type Plane struct {
	Rows int
	Columns int
}


func main() {
	fmt.Printf("hello")

	maxSeat := 0
	p := Plane{127,7}

	fileHandle, err := os.Open("day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)

	bp := BoardingPasses{
		 Seats: []Seat{},
	}

	for fileScanner.Scan() {
		s := fileScanner.Text()
		bp.AddSeat(Seat{
			Specification: s,
			SeatID: 0,
		})
	}

	for i:= 0; i<len(bp.Seats);i++{
		bp.Seats[i].CalcAndUpdateID(p)
		if bp.Seats[i].SeatID > maxSeat {
			maxSeat = bp.Seats[i].SeatID
		}
	}

	fmt.Printf("Max Seat ID: %d\n", maxSeat)

}
