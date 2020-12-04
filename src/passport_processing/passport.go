package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
		switch kvArray[0] {
		case "byr":
			f := Field{
				Key:       kvArray[0],
				Value:     kvArray[1],
				Validator: ByrValidator{},
			}
			p.Fields = append(p.Fields, f)
		case "iyr":
			f := Field{
				Key:       kvArray[0],
				Value:     kvArray[1],
				Validator: IyrValidator{},
			}
			p.Fields = append(p.Fields, f)
		case "eyr":
			f := Field{
				Key:       kvArray[0],
				Value:     kvArray[1],
				Validator: EyrValidator{},
			}
			p.Fields = append(p.Fields, f)
		case "hgt":
			f := Field{
				Key:       kvArray[0],
				Value:     kvArray[1],
				Validator: HgtValidator{},
			}
			p.Fields = append(p.Fields, f)
		case "hcl":
			f := Field{
				Key:       kvArray[0],
				Value:     kvArray[1],
				Validator: HclValidator{},
			}
			p.Fields = append(p.Fields, f)
		case "ecl":
			f := Field{
				Key:       kvArray[0],
				Value:     kvArray[1],
				Validator: EclValidator{},
			}
			p.Fields = append(p.Fields, f)
		case "pid":
			f := Field{
				Key:       kvArray[0],
				Value:     kvArray[1],
				Validator: PidValidator{},
			}
			p.Fields = append(p.Fields, f)
		default:
			f := Field{
				Key:       kvArray[0],
				Value:     kvArray[1],
				Validator: DefaultValidator{},
			}
			p.Fields = append(p.Fields, f)
		}
	}
	return p.Fields
}

func (p *Passport) IsValid() bool {
	valid := true
	for _, req := range MandatoryFields {
		found := false
		fieldIndex := 0
		for i, f := range p.Fields {
			if f.Key == req {
				found = true
				fieldIndex = i
				break
			}
		}
		if !found || !p.Fields[fieldIndex].Validator.IsValid(p.Fields[fieldIndex].Value) {
			valid = false
			break
		}
	}
	return valid
}

type FieldValidator interface {
	IsValid(value string) bool
}
type Field struct {
	Key       string
	Value     string
	Validator FieldValidator
}

var MandatoryFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

type DefaultValidator struct{}

func (dv DefaultValidator) IsValid(value string) bool {
	return true
}

type ByrValidator struct{}

func (v ByrValidator) IsValid(value string) bool {
	s, err := strconv.Atoi(value)
	if err == nil && s >= 1920 && s <= 2002 {
		return true
	}
	return false
}

type IyrValidator struct{}

func (v IyrValidator) IsValid(value string) bool {
	s, err := strconv.Atoi(value)
	if err == nil && s >= 2010 && s <= 2020 {
		return true
	}
	return false
}

type EyrValidator struct{}

func (v EyrValidator) IsValid(value string) bool {
	s, err := strconv.Atoi(value)
	if err == nil && s >= 2020 && s <= 2030 {
		return true
	}
	return false
}

type HgtValidator struct{}

func (v HgtValidator) IsValid(value string) bool {
	unit := value[len(value)-2:]
	s, err := strconv.Atoi(value[:len(value)-2])
	if unit == "cm" {
		if err == nil && s >= 150 && s <= 193 {
			return true
		}
	} else if unit == "in" {
		if err == nil && s >= 59 && s <= 76 {
			return true
		}
	}
	return false
}

type HclValidator struct{}

func (v HclValidator) IsValid(value string) bool {
	matched, err := regexp.MatchString(`^#[a-f0-9]{6}$`, value)
	if err == nil && matched {
		return true
	}
	return false
}

type EclValidator struct{}

func (v EclValidator) IsValid(value string) bool {
	valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, expected := range valid {
		if value == expected {
			return true
		}
	}
	return false
}

type PidValidator struct{}

func (v PidValidator) IsValid(value string) bool {
	matched, err := regexp.MatchString(`^[0-9]{9}$`, value)
	if err == nil && matched {
		return true
	}
	return false
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
