package day2

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//For a list of passwords in the format
//m-n a: s
//where m and n are integers such that m < n
//a is a letter a-z
//s is a string
//count how many passwords where s contains a number of a between m and n, inclusive
func Challenge1() (int, error) {
	var count int
	input, err := ioutil.ReadFile("./day2/d2.txt")
	if err != nil {
		return count, err
	}
	str := strings.TrimSpace(string(input))
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		var m, n int
		var a, s string
		_, err := fmt.Sscanf(line, "%d-%d %1s: %s", &m, &n, &a, &s)
		if err != nil {
			return count, err
		}
		occurrences := strings.Count(s, a)
		if m <= occurrences && occurrences <= n {
			count++
		}
	}

	return count, nil
}

//m or n must contain a, but not both
func Challenge2() (int, error) {
	var count int
	input, err := ioutil.ReadFile("./day2/d2.txt")
	if err != nil {
		return count, err
	}
	str := strings.TrimSpace(string(input))
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		var m, n int
		var a, s string
		var occurences = 0
		_, err := fmt.Sscanf(line, "%d-%d %1s: %s", &m, &n, &a, &s)
		if err != nil {
			return count, err
		}
		first := string(s[m-1])
		second := string(s[n-1])
		if first == a {
			occurences++
		}
		if second == a {
			occurences++
		}
		if occurences == 1 {
			count++
		}
	}

	return count, nil
}
