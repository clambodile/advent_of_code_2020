package day4

import (
	"github.com/clambodile/advent_of_code_2020/util/io"
	"regexp"
	"strconv"
	"strings"
)

func Challenge1() (int, error) {
	filename := "./day4/d4.txt"
	lineGroups, err := io.ReadLineGroups(filename)
	if err != nil {
		return 0, err
	}

	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	validCount := 0
	for _, lineGroup := range lineGroups {
		fieldRegex := regexp.MustCompile(`[a-zA-Z]+:\S`)
		fields := fieldRegex.FindAllString(lineGroup, -1)
		foundFields := map[string]bool{}
		for _, field := range fields {
			segments := strings.Split(field, ":")
			key := segments[0]
			foundFields[key] = true
		}
		valid := hasRequiredFields(requiredFields, foundFields)
		if valid {
			validCount++
		}
	}
	return validCount, nil
}

func Challenge2() (int, error) {
	filename := "./day4/d4.txt"
	lineGroups, err := io.ReadLineGroups(filename)
	if err != nil {
		return 0, err
	}

	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	fieldValidations := map[string]func(string) bool{
		"byr": func(birthYear string) bool {
			year, err := strconv.Atoi(birthYear)
			if err != nil {
				return false
			}
			return year >= 1920 && year <= 2002
		},
		"iyr": func(issueYear string) bool {
			year, err := strconv.Atoi(issueYear)
			if err != nil {
				return false
			}
			return year >= 2010 && year <= 2020
		},
		"eyr": func(expirationYear string) bool {
			year, err := strconv.Atoi(expirationYear)
			if err != nil {
				return false
			}
			return year >= 2020 && year <= 2030
		},
		"hgt": func(height string) bool {
			rgx := regexp.MustCompile(`(\d{2,3})(in|cm)`)
			matches := rgx.FindStringSubmatch(height)
			if len(matches) < 3 {
				return false
			}
			num, err := strconv.Atoi(matches[1])
			if err != nil {
				return false
			}
			unit := matches[2]
			if unit == "in" {
				return num >= 59 && num <= 76
			} else if unit == "cm" {
				return num >= 150 && num <= 193
			}
			return false
		},
		"hcl": func(hairColor string) bool {
			rgx := regexp.MustCompile(`#[0-9a-f]{6}`)
			return rgx.MatchString(hairColor)
		},
		"ecl": func(eyeColor string) bool {
			rgx := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
			return rgx.MatchString(eyeColor)
		},
		"pid": func(passportId string) bool {
			rgx := regexp.MustCompile(`^[0-9]{9}$`)
			return rgx.MatchString(passportId)
		},
		"cid": func(countryId string) bool {
			return true
		},
	}
	validCount := 0
	for _, lineGroup := range lineGroups {
		fieldRegex := regexp.MustCompile(`[a-zA-Z]+:\S+`)
		fields := fieldRegex.FindAllString(lineGroup, -1)
		foundFields := map[string]string{}
		existingFields := map[string]bool{}
		for _, field := range fields {
			segments := strings.Split(field, ":")
			key := segments[0]
			val := segments[1]
			foundFields[key] = val
			existingFields[key] = true
		}
		valid := hasRequiredFields(requiredFields, existingFields) && allFieldsValid(fieldValidations, foundFields)
		if valid {
			validCount++
		}
	}
	return validCount, nil
}

func allFieldsValid(fieldValidations map[string]func(string) bool, fields map[string]string) bool {
	for k, v := range fields {
		f := fieldValidations[k]
		if !f(v) {
			return false
		}
	}
	return true
}

func hasRequiredFields(requiredFields []string, mp map[string]bool) bool {
	for _, field := range requiredFields {
		_, ok := mp[field]
		if !ok {
			return false
		}
	}
	return true
}
