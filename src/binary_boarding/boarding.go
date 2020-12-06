package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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
	Row int
	Col	int
	SeatID	int
}

func (s *Seat) CalcAndLocateSeat(p Plane) {
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
	s.Row = endingRow
	s.Col = endingCol
	s.SeatID = endingRow * 8 + endingCol
}


type Plane struct {
	Rows int
	Columns int
}

func (p *Plane) FindEmptySeats(b BoardingPasses) []Seat {
	
	emptySeats := []Seat{}
	for i:=0; i<= p.Rows; i++ {
		for j:=0; j <= p.Columns; j++ {
			found := false
			for _, s := range b.Seats {
				if s.Row == i && s.Col == j {
					found = true
				}
			}
			if found == false {
				emptySeats = append(emptySeats, Seat{
					Specification: "",
					SeatID: i * 8 + j,
					Row: i,
					Col: j,
				})
			}
		}
	}
	return emptySeats
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

	b := BoardingPasses{
		 Seats: []Seat{},
	}

	for fileScanner.Scan() {
		s := fileScanner.Text()
		b.AddSeat(Seat{
			Specification: s,
			SeatID: 0,
		})
	}
    start := time.Now()

	for i:= 0; i<len(b.Seats);i++{
		b.Seats[i].CalcAndLocateSeat(p)
		fmt.Printf("Seat ID: %d\n", b.Seats[i].SeatID)

		if b.Seats[i].SeatID > maxSeat {
			maxSeat = b.Seats[i].SeatID
		}
	}

	fmt.Printf("Max Seat ID: %d\n", maxSeat)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	
	emptySeats := p.FindEmptySeats(b)
	for i:= 0; i< len(emptySeats); i++ {
		fmt.Printf("Empty Seat Row: %d, Col: %d, SeatID: %d\n", emptySeats[i].Row,emptySeats[i].Col,emptySeats[i].SeatID)

	}

	elapsed = time.Since(start)
    log.Printf("Binomial took %s", elapsed)

}
