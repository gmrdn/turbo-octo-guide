package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var MandatoryFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

type BatchFile struct {
	Passports []Passport
}

func (b *BatchFile) AddPassport(p Passport) []Passport {
	b.Passports = append(b.Passports, p)
	return b.Passports
}

type Passport struct {
	Fields []Field
}

func (p *Passport) AddFields(s string) []Field {
	keyValues := strings.Split(s, " ")
	for _, kv := range keyValues {
		kvArray := strings.Split(kv, ":")
		f := Field{
			Key:   kvArray[0],
			Value: kvArray[1],
		}
		p.Fields = append(p.Fields, f)
	}
	return p.Fields
}

func (p *Passport) IsValid() bool {
	valid := true
	for _, req := range MandatoryFields {
		found := false
		for _, f := range p.Fields {
			if f.Key == req {
				found = true
				break
			}
		}
		if found == false {
			valid = false
			break
		}
	}
	return valid
}

type Field struct {
	Key   string
	Value string
}

func main() {
	fmt.Printf("hello")

	fileHandle, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)

	b := BatchFile{
		Passports: []Passport{},
	}

	b.AddPassport(Passport{
		Fields: []Field{},
	})

	currentIndex := 0

	for fileScanner.Scan() {
		s := fileScanner.Text()
		if s == "" {
			b.AddPassport(Passport{
				Fields: []Field{},
			})
			currentIndex++
		} else {
			b.Passports[currentIndex].AddFields(s)
		}
	}

	nbValid := 0
	for _, p := range b.Passports {
		if p.IsValid() {
			nbValid++
		}
	}

	fmt.Printf("Total: %d\n", nbValid)

}
