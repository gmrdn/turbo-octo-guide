package main

import "testing"

func InitBatchFile() BatchFile {
	b := BatchFile{
		Passports: []Passport{},
	}
	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.Passports[0].AddFields("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd")
	b.Passports[0].AddFields("byr:1937 iyr:2017 cid:147 hgt:183cm")

	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.Passports[1].AddFields("iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884")
	b.Passports[1].AddFields("hcl:#cfa07d byr:1929")

	b.AddPassport(Passport{
		Fields: []Field{},
	})

	b.Passports[2].AddFields("hcl:#ae17e1 iyr:2013")
	b.Passports[2].AddFields("eyr:2024")
	b.Passports[2].AddFields("ecl:brn pid:760753108 byr:1931")
	b.Passports[2].AddFields("hgt:179cm")

	b.AddPassport(Passport{
		Fields: []Field{},
	})

	b.Passports[3].AddFields("hcl:#cfa07d eyr:2025 pid:166559648")
	b.Passports[3].AddFields("iyr:2011 ecl:brn hgt:59in")

	return b
}

func TestIsValid(t *testing.T) {

	b := InitBatchFile()

	want := true
	got := b.Passports[0].IsValid()
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}

func TestIsNotValid(t *testing.T) {

	b := InitBatchFile()

	want := false
	got := b.Passports[1].IsValid()
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}

func TestIsValidEvenWithoutCID(t *testing.T) {

	b := InitBatchFile()

	want := true
	got := b.Passports[2].IsValid()
	if got != want {
		t.Errorf("got: %t, instead of: %t.", got, want)
	}
}
