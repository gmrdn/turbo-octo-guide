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

func InitInvalidPassowrds() BatchFile {
	b := BatchFile{
		Passports: []Passport{},
	}
	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.Passports[0].AddFields("eyr:1972 cid:100")
	b.Passports[0].AddFields("hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926")
	b.Passports[1].AddFields("iyr:2019")
	b.Passports[1].AddFields("hcl:#602927 eyr:1967 hgt:170cm")
	b.Passports[1].AddFields("ecl:grn pid:012533040 byr:1946")
	b.Passports[2].AddFields("hcl:dab227 iyr:2012")
	b.Passports[2].AddFields("ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277")
	b.Passports[3].AddFields("hgt:59cm ecl:zzz")
	b.Passports[3].AddFields("eyr:2038 hcl:74454a iyr:2023")
	b.Passports[3].AddFields("pid:3556412378 byr:2007")

	return b
}

func InitValidPassowrds() BatchFile {
	b := BatchFile{
		Passports: []Passport{},
	}
	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.AddPassport(Passport{
		Fields: []Field{},
	})
	b.AddPassport(Passport{
		Fields: []Field{},
	})

	b.Passports[0].AddFields("pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980")
	b.Passports[0].AddFields("hcl:#623a2f")
	b.Passports[1].AddFields("eyr:2029 ecl:blu cid:129 byr:1989")
	b.Passports[1].AddFields("iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm")
	b.Passports[2].AddFields("hcl:#888785")
	b.Passports[2].AddFields("hgt:164cm byr:2001 iyr:2015 cid:88")
	b.Passports[2].AddFields("pid:545766238 ecl:hzl")
	b.Passports[2].AddFields("eyr:2022")
	b.Passports[3].AddFields("iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719")
	return b
}

func TestAreAllInvalid(t *testing.T) {
	b := InitInvalidPassowrds()

	want := false
	for _, p := range b.Passports {
		got := p.IsValid()
		if got != want {
			t.Errorf("got: %t, instead of: %t.", got, want)
		}
	}
}

func TestAreAllValid(t *testing.T) {
	b := InitValidPassowrds()

	want := true
	for _, p := range b.Passports {
		got := p.IsValid()
		if got != want {
			t.Errorf("got: %t, instead of: %t.", got, want)
		}
	}
}
