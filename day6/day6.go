package day6

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func Challenge1() (int, error) {
	filename := "./day6/d6.txt"
	answerGroups, err := readAnswerGroupsAny(filename)
	if err != nil {
		return 0, err
	}
	totalAnswers := 0
	for _, group := range answerGroups {
		totalAnswers += answerTotal(group)
	}
	return totalAnswers, nil
}

func Challenge2() (int, error) {
	filename := "./day6/d6.txt"
	answerGroups, err := readAnswerGroupsAll(filename)
	if err != nil {
		return 0, err
	}
	totalAnswers := 0
	for _, group := range answerGroups {
		totalAnswers += answerTotal(group)
	}
	return totalAnswers, nil
}

func answerTotal(answerGroup map[string]bool) int {
	total := 0
	for _, v := range answerGroup {
		if v {
			total++
		}
	}
	return total
}

func readAnswerGroupsAny(filename string) ([]map[string]bool, error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	answerGroups := strings.Split(string(input), "\n\n")
	output := make([]map[string]bool, len(answerGroups))
	for i, group := range answerGroups {
		output[i] = readAnswerGroupAny(group)
	}
	return output, nil
}

func readAnswerGroupsAll(filename string) ([]map[string]bool, error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	answerGroups := strings.Split(string(input), "\n\n")
	output := make([]map[string]bool, len(answerGroups))
	for i, group := range answerGroups {
		output[i] = readAnswerGroupAll(group)
	}
	return output, nil
}

//Returns an answer group where each key corresponds to an answer to which someone responded "true".
func readAnswerGroupAny(group string) map[string]bool {
	lines := strings.Split(group, "\n")
	answerGroup := make(map[string]bool)
	for _, line := range lines {
		for _, rn := range line {
			char := string(rn)
			answerGroup[char] = true
		}
	}
	return answerGroup
}

//Returns an answer group where each key corresponds to an answer to which all respondents answered "true".
func readAnswerGroupAll(group string) map[string]bool {
	lines := strings.Split(group, "\n")
	answerGroup := make(map[string]bool)
	for i, line := range lines {
		if i == 0 {
			for _, rn := range line {
				char := string(rn)
				answerGroup[char] = true
			}
			continue
		}
		for k, v := range answerGroup {
			if v {
				matches, err := regexp.MatchString(k, line)
				if err != nil {
					continue
				}
				answerGroup[k] = matches
			}
		}
	}
	return answerGroup
}
